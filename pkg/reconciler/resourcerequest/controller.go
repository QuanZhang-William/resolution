/*
Copyright 2022 The Tekton Authors

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package resourcerequest

import (
	"context"

	"knative.dev/pkg/configmap"
	"knative.dev/pkg/controller"

	resourcerequestinformer "github.com/tektoncd/resolution/pkg/client/injection/informers/resolution/v1alpha1/resourcerequest"
	resourcerequestreconciler "github.com/tektoncd/resolution/pkg/client/injection/reconciler/resolution/v1alpha1/resourcerequest"
)

func NewController(ctx context.Context, cmw configmap.Watcher) *controller.Impl {
	r := &Reconciler{}
	impl := resourcerequestreconciler.NewImpl(ctx, r)

	reqinformer := resourcerequestinformer.Get(ctx)
	reqinformer.Informer().AddEventHandler(controller.HandleAll(impl.Enqueue))

	return impl
}