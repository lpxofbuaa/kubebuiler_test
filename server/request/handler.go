package request

import (
	"context"
	"errors"

	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	crdtryv1 "kubebuilder.test/crdtry/api/crdtry/v1"
	crdtryVersioned "kubebuilder.test/crdtry/generated/crdtry/clientset/versioned"
)

type Handler struct {
	Clientset *crdtryVersioned.Clientset
}

func (h *Handler) HandleCreateLpxpodReq(req *LpxpodRequest) (*crdtryv1.Lpxpod, error) {
	if req.Name == "" {
		return nil, errors.New("Name cannot be empty!")
	}
	if req.Namespace == "" {
		req.Namespace = "default"
	}
	lpxpod := &crdtryv1.Lpxpod{
		ObjectMeta: v1.ObjectMeta{
			Name:      req.Name,
			Namespace: req.Namespace,
		},
		Spec: crdtryv1.LpxpodSpec{
			Podport:        req.Podport,
			Deploymentname: req.Deploymentname,
			Schedule:       req.Schedule,
		},
	}
	lpxpod, err := h.Clientset.CrdtryV1().Lpxpods(req.Namespace).Create(context.TODO(), lpxpod, v1.CreateOptions{})
	return lpxpod, err
}

func (h *Handler) HandleGetLpxpodReq(req *LpxpodRequest) (*crdtryv1.Lpxpod, error) {
	if req.Name == "" {
		return nil, errors.New("Name cannot be empty!")
	}
	if req.Namespace == "" {
		req.Namespace = "default"
	}
	lpxpod, err := h.Clientset.CrdtryV1().Lpxpods(req.Namespace).Get(context.TODO(), req.Name, v1.GetOptions{})
	return lpxpod, err
}
