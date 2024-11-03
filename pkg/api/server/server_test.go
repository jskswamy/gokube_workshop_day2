package server

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/emicklei/go-restful/v3"
	"github.com/stretchr/testify/assert"
	clientv3 "go.etcd.io/etcd/client/v3"
	"gokube/pkg/api"
	"gokube/pkg/storage"
)

func TestCreateNode(t *testing.T) {
	storage.TestWithEmbeddedEtcd(t, func(t *testing.T, etcdServer *clientv3.Client) {
		store := storage.NewEtcdStorage(etcdServer)
		apiServer := NewAPIServer(store)

		node := &api.Node{
			ObjectMeta: api.ObjectMeta{
				Name: "test-node",
			},
			Status: api.NodeReady,
		}

		body, _ := json.Marshal(node)
		req := httptest.NewRequest("POST", "/api/v1/nodes", bytes.NewReader(body))
		req.Header.Set("Content-Type", restful.MIME_JSON)
		resp := httptest.NewRecorder()

		container := restful.NewContainer()
		apiServer.registerRoutes(container)
		container.ServeHTTP(resp, req)

		assert.Equal(t, http.StatusCreated, resp.Code)

		var createdNode api.Node
		err := json.Unmarshal(resp.Body.Bytes(), &createdNode)
		assert.NoError(t, err)
		assert.Equal(t, node.Name, createdNode.Name)
		assert.Equal(t, node.Status, createdNode.Status)
	})
}

func TestUpdateNodeStatus(t *testing.T) {
	storage.TestWithEmbeddedEtcd(t, func(t *testing.T, etcdServer *clientv3.Client) {
		store := storage.NewEtcdStorage(etcdServer)
		apiServer := NewAPIServer(store)

		// First, create a node
		node := &api.Node{
			ObjectMeta: api.ObjectMeta{
				Name: "test-node",
			},
			Status: api.NodeReady,
		}

		err := apiServer.nodeRegistry.CreateNode(context.Background(), node)
		assert.NoError(t, err)

		// Now, update the node's status
		updatedNode := &api.Node{
			ObjectMeta: api.ObjectMeta{
				Name: "test-node",
			},
			Status: api.NodeNotReady,
		}

		body, _ := json.Marshal(updatedNode)
		req := httptest.NewRequest("PUT", "/api/v1/nodes/test-node", bytes.NewReader(body))
		req.Header.Set("Content-Type", restful.MIME_JSON)
		resp := httptest.NewRecorder()

		container := restful.NewContainer()
		apiServer.registerRoutes(container)
		container.ServeHTTP(resp, req)

		assert.Equal(t, http.StatusOK, resp.Code)

		var returnedNode api.Node
		err = json.Unmarshal(resp.Body.Bytes(), &returnedNode)
		assert.NoError(t, err)
		assert.Equal(t, updatedNode.Name, returnedNode.Name)
		assert.Equal(t, updatedNode.Status, returnedNode.Status)
	})
}

func TestCreatePod(t *testing.T) {
	storage.TestWithEmbeddedEtcd(t, func(t *testing.T, etcdServer *clientv3.Client) {
		store := storage.NewEtcdStorage(etcdServer)
		apiServer := NewAPIServer(store)

		pod := &api.Pod{
			ObjectMeta: api.ObjectMeta{
				Name: "test-pod",
			},
			Spec: api.PodSpec{
				Replicas: 1,
				Containers: []api.Container{
					{
						Image: "nginx:latest",
					},
				},
			},
			// Note: We don't set the Status field here, as it should be set by the server
		}

		body, _ := json.Marshal(pod)
		req := httptest.NewRequest("POST", "/api/v1/pods", bytes.NewReader(body))
		req.Header.Set("Content-Type", restful.MIME_JSON)
		resp := httptest.NewRecorder()

		container := restful.NewContainer()
		apiServer.registerRoutes(container)
		container.ServeHTTP(resp, req)

		assert.Equal(t, http.StatusCreated, resp.Code)

		var createdPod api.Pod
		err := json.Unmarshal(resp.Body.Bytes(), &createdPod)
		assert.NoError(t, err)
		assert.Equal(t, pod.Name, createdPod.Name)
		assert.Equal(t, pod.Spec.Replicas, createdPod.Spec.Replicas)
		assert.Equal(t, len(pod.Spec.Containers), len(createdPod.Spec.Containers))
		assert.Equal(t, pod.Spec.Containers[0].Image, createdPod.Spec.Containers[0].Image)

		// Check that the status is set to Unassigned
		assert.Equal(t, api.PodPending, createdPod.Status)
	})
}

func TestUpdatePod(t *testing.T) {
	updatePodStatus := func(apiServerURL string, pod *api.Pod) error {
		url := fmt.Sprintf("http://%s/api/v1/pods/%s/", apiServerURL, pod.Name)

		jsonData, err := json.Marshal(pod)
		if err != nil {
			return fmt.Errorf("failed to marshal pod data: %w", err)
		}

		req, err := http.NewRequest("PUT", url, bytes.NewBuffer(jsonData))
		if err != nil {
			return fmt.Errorf("failed to create request: %w", err)
		}
		req.Header.Set("Content-Type", restful.MIME_JSON)

		resp, err := http.DefaultClient.Do(req)
		if err != nil {
			return fmt.Errorf("failed to send request to API server: %w", err)
		}
		defer resp.Body.Close()

		if resp.StatusCode != http.StatusOK {
			return fmt.Errorf("failed to update pod status, status code: %d", resp.StatusCode)
		}

		log.Printf("Updated pod status for %s: %v", pod.Name, pod.Status)

		return nil
	}

	storage.TestWithEmbeddedEtcd(t, func(t *testing.T, etcdServer *clientv3.Client) {
		store := storage.NewEtcdStorage(etcdServer)
		apiServer := NewAPIServer(store)

		go func() {
			err := apiServer.Start("localhost:8080")
			if err != nil {
				log.Fatalf("Failed to start API server: %v", err)
			}
		}()

		// Setup
		pod := &api.Pod{
			ObjectMeta: api.ObjectMeta{
				Name: "test-pod",
			},
			Spec: api.PodSpec{
				Replicas: 1,
				Containers: []api.Container{
					{
						Image: "nginx:latest",
					},
				},
			},
			// Note: We don't set the Status field here, as it should be set by the server
		}

		body, _ := json.Marshal(pod)
		req, _ := http.NewRequest("POST", "http://localhost:8080/api/v1/pods", bytes.NewReader(body))
		req.Header.Set("Content-Type", restful.MIME_JSON)

		resp, err := http.DefaultClient.Do(req)
		assert.NoError(t, err)

		assert.Equal(t, http.StatusCreated, resp.StatusCode)

		var createdPod api.Pod
		body, _ = io.ReadAll(resp.Body)
		err = json.Unmarshal(body, &createdPod)
		assert.NoError(t, err)
		assert.Equal(t, pod.Name, createdPod.Name)
		assert.Equal(t, pod.Spec.Replicas, createdPod.Spec.Replicas)
		assert.Equal(t, len(pod.Spec.Containers), len(createdPod.Spec.Containers))
		assert.Equal(t, pod.Spec.Containers[0].Image, createdPod.Spec.Containers[0].Image)

		// Check that the status is set to Unassigned
		assert.Equal(t, api.PodPending, createdPod.Status)

		// Update the pod status
		pod.Status = api.PodRunning

		err = updatePodStatus("localhost:8080", pod)
		assert.NoError(t, err)
	})
}
