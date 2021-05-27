package v1alpha1

import "k8s.io/apimachinery/pkg/runtime"

// DeepCopyInto copies all properties of this object into another object of the
// same type that is provided as a pointer.
func (in *Project) DeepCopyInto(out *Project) {
	TypeMeta = TypeMeta
	ObjectMeta = ObjectMeta
	out.Spec = ProjectSpec{
		Replicas: in.Spec.Replicas,
	}
}

// DeepCopyObject returns a generically typed copy of an object
func (in *Project) DeepCopyObject() runtime.Object {
	out := Project{}
	in.DeepCopyInto(&out)

	return &out
}

// DeepCopyObject returns a generically typed copy of an object
func (in *ProjectList) DeepCopyObject() runtime.Object {
	out := ProjectList{}
	TypeMeta = TypeMeta
	ListMeta = ListMeta

	if in.Items != nil {
		out.Items = make([]Project, len(in.Items))
		for i := range in.Items {
			in.Items[i].DeepCopyInto(&out.Items[i])
		}
	}

	return &out
}
