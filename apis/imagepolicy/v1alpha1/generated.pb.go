// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: k8s.io/api/imagepolicy/v1alpha1/generated.proto

/*
	Package v1alpha1 is a generated protocol buffer package.

	It is generated from these files:
		k8s.io/api/imagepolicy/v1alpha1/generated.proto

	It has these top-level messages:
		ImageReview
		ImageReviewContainerSpec
		ImageReviewSpec
		ImageReviewStatus
*/
package v1alpha1

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import k8s_io_apimachinery_pkg_apis_meta_v1 "github.com/tryggth/k8s/apis/meta/v1"
import _ "github.com/tryggth/k8s/runtime"
import _ "github.com/tryggth/k8s/runtime/schema"
import _ "github.com/tryggth/k8s/util/intstr"

import io "io"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

// ImageReview checks if the set of images in a pod are allowed.
type ImageReview struct {
	// +optional
	Metadata *k8s_io_apimachinery_pkg_apis_meta_v1.ObjectMeta `protobuf:"bytes,1,opt,name=metadata" json:"metadata,omitempty"`
	// Spec holds information about the pod being evaluated
	Spec *ImageReviewSpec `protobuf:"bytes,2,opt,name=spec" json:"spec,omitempty"`
	// Status is filled in by the backend and indicates whether the pod should be allowed.
	// +optional
	Status           *ImageReviewStatus `protobuf:"bytes,3,opt,name=status" json:"status,omitempty"`
	XXX_unrecognized []byte             `json:"-"`
}

func (m *ImageReview) Reset()                    { *m = ImageReview{} }
func (m *ImageReview) String() string            { return proto.CompactTextString(m) }
func (*ImageReview) ProtoMessage()               {}
func (*ImageReview) Descriptor() ([]byte, []int) { return fileDescriptorGenerated, []int{0} }

func (m *ImageReview) GetMetadata() *k8s_io_apimachinery_pkg_apis_meta_v1.ObjectMeta {
	if m != nil {
		return m.Metadata
	}
	return nil
}

func (m *ImageReview) GetSpec() *ImageReviewSpec {
	if m != nil {
		return m.Spec
	}
	return nil
}

func (m *ImageReview) GetStatus() *ImageReviewStatus {
	if m != nil {
		return m.Status
	}
	return nil
}

// ImageReviewContainerSpec is a description of a container within the pod creation request.
type ImageReviewContainerSpec struct {
	// This can be in the form image:tag or image@SHA:012345679abcdef.
	// +optional
	Image            *string `protobuf:"bytes,1,opt,name=image" json:"image,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *ImageReviewContainerSpec) Reset()         { *m = ImageReviewContainerSpec{} }
func (m *ImageReviewContainerSpec) String() string { return proto.CompactTextString(m) }
func (*ImageReviewContainerSpec) ProtoMessage()    {}
func (*ImageReviewContainerSpec) Descriptor() ([]byte, []int) {
	return fileDescriptorGenerated, []int{1}
}

func (m *ImageReviewContainerSpec) GetImage() string {
	if m != nil && m.Image != nil {
		return *m.Image
	}
	return ""
}

// ImageReviewSpec is a description of the pod creation request.
type ImageReviewSpec struct {
	// Containers is a list of a subset of the information in each container of the Pod being created.
	// +optional
	Containers []*ImageReviewContainerSpec `protobuf:"bytes,1,rep,name=containers" json:"containers,omitempty"`
	// Annotations is a list of key-value pairs extracted from the Pod's annotations.
	// It only includes keys which match the pattern `*.image-policy.k8s.io/*`.
	// It is up to each webhook backend to determine how to interpret these annotations, if at all.
	// +optional
	Annotations map[string]string `protobuf:"bytes,2,rep,name=annotations" json:"annotations,omitempty" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"bytes,2,opt,name=value"`
	// Namespace is the namespace the pod is being created in.
	// +optional
	Namespace        *string `protobuf:"bytes,3,opt,name=namespace" json:"namespace,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *ImageReviewSpec) Reset()                    { *m = ImageReviewSpec{} }
func (m *ImageReviewSpec) String() string            { return proto.CompactTextString(m) }
func (*ImageReviewSpec) ProtoMessage()               {}
func (*ImageReviewSpec) Descriptor() ([]byte, []int) { return fileDescriptorGenerated, []int{2} }

func (m *ImageReviewSpec) GetContainers() []*ImageReviewContainerSpec {
	if m != nil {
		return m.Containers
	}
	return nil
}

func (m *ImageReviewSpec) GetAnnotations() map[string]string {
	if m != nil {
		return m.Annotations
	}
	return nil
}

func (m *ImageReviewSpec) GetNamespace() string {
	if m != nil && m.Namespace != nil {
		return *m.Namespace
	}
	return ""
}

// ImageReviewStatus is the result of the token authentication request.
type ImageReviewStatus struct {
	// Allowed indicates that all images were allowed to be run.
	Allowed *bool `protobuf:"varint,1,opt,name=allowed" json:"allowed,omitempty"`
	// Reason should be empty unless Allowed is false in which case it
	// may contain a short description of what is wrong.  Kubernetes
	// may truncate excessively long errors when displaying to the user.
	// +optional
	Reason           *string `protobuf:"bytes,2,opt,name=reason" json:"reason,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *ImageReviewStatus) Reset()                    { *m = ImageReviewStatus{} }
func (m *ImageReviewStatus) String() string            { return proto.CompactTextString(m) }
func (*ImageReviewStatus) ProtoMessage()               {}
func (*ImageReviewStatus) Descriptor() ([]byte, []int) { return fileDescriptorGenerated, []int{3} }

func (m *ImageReviewStatus) GetAllowed() bool {
	if m != nil && m.Allowed != nil {
		return *m.Allowed
	}
	return false
}

func (m *ImageReviewStatus) GetReason() string {
	if m != nil && m.Reason != nil {
		return *m.Reason
	}
	return ""
}

func init() {
	proto.RegisterType((*ImageReview)(nil), "k8s.io.api.imagepolicy.v1alpha1.ImageReview")
	proto.RegisterType((*ImageReviewContainerSpec)(nil), "k8s.io.api.imagepolicy.v1alpha1.ImageReviewContainerSpec")
	proto.RegisterType((*ImageReviewSpec)(nil), "k8s.io.api.imagepolicy.v1alpha1.ImageReviewSpec")
	proto.RegisterType((*ImageReviewStatus)(nil), "k8s.io.api.imagepolicy.v1alpha1.ImageReviewStatus")
}
func (m *ImageReview) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *ImageReview) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if m.Metadata != nil {
		dAtA[i] = 0xa
		i++
		i = encodeVarintGenerated(dAtA, i, uint64(m.Metadata.Size()))
		n1, err := m.Metadata.MarshalTo(dAtA[i:])
		if err != nil {
			return 0, err
		}
		i += n1
	}
	if m.Spec != nil {
		dAtA[i] = 0x12
		i++
		i = encodeVarintGenerated(dAtA, i, uint64(m.Spec.Size()))
		n2, err := m.Spec.MarshalTo(dAtA[i:])
		if err != nil {
			return 0, err
		}
		i += n2
	}
	if m.Status != nil {
		dAtA[i] = 0x1a
		i++
		i = encodeVarintGenerated(dAtA, i, uint64(m.Status.Size()))
		n3, err := m.Status.MarshalTo(dAtA[i:])
		if err != nil {
			return 0, err
		}
		i += n3
	}
	if m.XXX_unrecognized != nil {
		i += copy(dAtA[i:], m.XXX_unrecognized)
	}
	return i, nil
}

func (m *ImageReviewContainerSpec) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *ImageReviewContainerSpec) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if m.Image != nil {
		dAtA[i] = 0xa
		i++
		i = encodeVarintGenerated(dAtA, i, uint64(len(*m.Image)))
		i += copy(dAtA[i:], *m.Image)
	}
	if m.XXX_unrecognized != nil {
		i += copy(dAtA[i:], m.XXX_unrecognized)
	}
	return i, nil
}

func (m *ImageReviewSpec) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *ImageReviewSpec) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if len(m.Containers) > 0 {
		for _, msg := range m.Containers {
			dAtA[i] = 0xa
			i++
			i = encodeVarintGenerated(dAtA, i, uint64(msg.Size()))
			n, err := msg.MarshalTo(dAtA[i:])
			if err != nil {
				return 0, err
			}
			i += n
		}
	}
	if len(m.Annotations) > 0 {
		for k, _ := range m.Annotations {
			dAtA[i] = 0x12
			i++
			v := m.Annotations[k]
			mapSize := 1 + len(k) + sovGenerated(uint64(len(k))) + 1 + len(v) + sovGenerated(uint64(len(v)))
			i = encodeVarintGenerated(dAtA, i, uint64(mapSize))
			dAtA[i] = 0xa
			i++
			i = encodeVarintGenerated(dAtA, i, uint64(len(k)))
			i += copy(dAtA[i:], k)
			dAtA[i] = 0x12
			i++
			i = encodeVarintGenerated(dAtA, i, uint64(len(v)))
			i += copy(dAtA[i:], v)
		}
	}
	if m.Namespace != nil {
		dAtA[i] = 0x1a
		i++
		i = encodeVarintGenerated(dAtA, i, uint64(len(*m.Namespace)))
		i += copy(dAtA[i:], *m.Namespace)
	}
	if m.XXX_unrecognized != nil {
		i += copy(dAtA[i:], m.XXX_unrecognized)
	}
	return i, nil
}

func (m *ImageReviewStatus) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *ImageReviewStatus) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if m.Allowed != nil {
		dAtA[i] = 0x8
		i++
		if *m.Allowed {
			dAtA[i] = 1
		} else {
			dAtA[i] = 0
		}
		i++
	}
	if m.Reason != nil {
		dAtA[i] = 0x12
		i++
		i = encodeVarintGenerated(dAtA, i, uint64(len(*m.Reason)))
		i += copy(dAtA[i:], *m.Reason)
	}
	if m.XXX_unrecognized != nil {
		i += copy(dAtA[i:], m.XXX_unrecognized)
	}
	return i, nil
}

func encodeVarintGenerated(dAtA []byte, offset int, v uint64) int {
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return offset + 1
}
func (m *ImageReview) Size() (n int) {
	var l int
	_ = l
	if m.Metadata != nil {
		l = m.Metadata.Size()
		n += 1 + l + sovGenerated(uint64(l))
	}
	if m.Spec != nil {
		l = m.Spec.Size()
		n += 1 + l + sovGenerated(uint64(l))
	}
	if m.Status != nil {
		l = m.Status.Size()
		n += 1 + l + sovGenerated(uint64(l))
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func (m *ImageReviewContainerSpec) Size() (n int) {
	var l int
	_ = l
	if m.Image != nil {
		l = len(*m.Image)
		n += 1 + l + sovGenerated(uint64(l))
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func (m *ImageReviewSpec) Size() (n int) {
	var l int
	_ = l
	if len(m.Containers) > 0 {
		for _, e := range m.Containers {
			l = e.Size()
			n += 1 + l + sovGenerated(uint64(l))
		}
	}
	if len(m.Annotations) > 0 {
		for k, v := range m.Annotations {
			_ = k
			_ = v
			mapEntrySize := 1 + len(k) + sovGenerated(uint64(len(k))) + 1 + len(v) + sovGenerated(uint64(len(v)))
			n += mapEntrySize + 1 + sovGenerated(uint64(mapEntrySize))
		}
	}
	if m.Namespace != nil {
		l = len(*m.Namespace)
		n += 1 + l + sovGenerated(uint64(l))
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func (m *ImageReviewStatus) Size() (n int) {
	var l int
	_ = l
	if m.Allowed != nil {
		n += 2
	}
	if m.Reason != nil {
		l = len(*m.Reason)
		n += 1 + l + sovGenerated(uint64(l))
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func sovGenerated(x uint64) (n int) {
	for {
		n++
		x >>= 7
		if x == 0 {
			break
		}
	}
	return n
}
func sozGenerated(x uint64) (n int) {
	return sovGenerated(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *ImageReview) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowGenerated
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: ImageReview: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: ImageReview: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Metadata", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenerated
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthGenerated
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Metadata == nil {
				m.Metadata = &k8s_io_apimachinery_pkg_apis_meta_v1.ObjectMeta{}
			}
			if err := m.Metadata.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Spec", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenerated
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthGenerated
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Spec == nil {
				m.Spec = &ImageReviewSpec{}
			}
			if err := m.Spec.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Status", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenerated
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthGenerated
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Status == nil {
				m.Status = &ImageReviewStatus{}
			}
			if err := m.Status.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipGenerated(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthGenerated
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			m.XXX_unrecognized = append(m.XXX_unrecognized, dAtA[iNdEx:iNdEx+skippy]...)
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *ImageReviewContainerSpec) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowGenerated
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: ImageReviewContainerSpec: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: ImageReviewContainerSpec: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Image", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenerated
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthGenerated
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			s := string(dAtA[iNdEx:postIndex])
			m.Image = &s
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipGenerated(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthGenerated
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			m.XXX_unrecognized = append(m.XXX_unrecognized, dAtA[iNdEx:iNdEx+skippy]...)
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *ImageReviewSpec) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowGenerated
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: ImageReviewSpec: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: ImageReviewSpec: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Containers", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenerated
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthGenerated
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Containers = append(m.Containers, &ImageReviewContainerSpec{})
			if err := m.Containers[len(m.Containers)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Annotations", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenerated
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthGenerated
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Annotations == nil {
				m.Annotations = make(map[string]string)
			}
			var mapkey string
			var mapvalue string
			for iNdEx < postIndex {
				entryPreIndex := iNdEx
				var wire uint64
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return ErrIntOverflowGenerated
					}
					if iNdEx >= l {
						return io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					wire |= (uint64(b) & 0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				fieldNum := int32(wire >> 3)
				if fieldNum == 1 {
					var stringLenmapkey uint64
					for shift := uint(0); ; shift += 7 {
						if shift >= 64 {
							return ErrIntOverflowGenerated
						}
						if iNdEx >= l {
							return io.ErrUnexpectedEOF
						}
						b := dAtA[iNdEx]
						iNdEx++
						stringLenmapkey |= (uint64(b) & 0x7F) << shift
						if b < 0x80 {
							break
						}
					}
					intStringLenmapkey := int(stringLenmapkey)
					if intStringLenmapkey < 0 {
						return ErrInvalidLengthGenerated
					}
					postStringIndexmapkey := iNdEx + intStringLenmapkey
					if postStringIndexmapkey > l {
						return io.ErrUnexpectedEOF
					}
					mapkey = string(dAtA[iNdEx:postStringIndexmapkey])
					iNdEx = postStringIndexmapkey
				} else if fieldNum == 2 {
					var stringLenmapvalue uint64
					for shift := uint(0); ; shift += 7 {
						if shift >= 64 {
							return ErrIntOverflowGenerated
						}
						if iNdEx >= l {
							return io.ErrUnexpectedEOF
						}
						b := dAtA[iNdEx]
						iNdEx++
						stringLenmapvalue |= (uint64(b) & 0x7F) << shift
						if b < 0x80 {
							break
						}
					}
					intStringLenmapvalue := int(stringLenmapvalue)
					if intStringLenmapvalue < 0 {
						return ErrInvalidLengthGenerated
					}
					postStringIndexmapvalue := iNdEx + intStringLenmapvalue
					if postStringIndexmapvalue > l {
						return io.ErrUnexpectedEOF
					}
					mapvalue = string(dAtA[iNdEx:postStringIndexmapvalue])
					iNdEx = postStringIndexmapvalue
				} else {
					iNdEx = entryPreIndex
					skippy, err := skipGenerated(dAtA[iNdEx:])
					if err != nil {
						return err
					}
					if skippy < 0 {
						return ErrInvalidLengthGenerated
					}
					if (iNdEx + skippy) > postIndex {
						return io.ErrUnexpectedEOF
					}
					iNdEx += skippy
				}
			}
			m.Annotations[mapkey] = mapvalue
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Namespace", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenerated
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthGenerated
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			s := string(dAtA[iNdEx:postIndex])
			m.Namespace = &s
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipGenerated(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthGenerated
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			m.XXX_unrecognized = append(m.XXX_unrecognized, dAtA[iNdEx:iNdEx+skippy]...)
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *ImageReviewStatus) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowGenerated
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: ImageReviewStatus: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: ImageReviewStatus: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Allowed", wireType)
			}
			var v int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenerated
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				v |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			b := bool(v != 0)
			m.Allowed = &b
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Reason", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenerated
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthGenerated
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			s := string(dAtA[iNdEx:postIndex])
			m.Reason = &s
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipGenerated(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthGenerated
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			m.XXX_unrecognized = append(m.XXX_unrecognized, dAtA[iNdEx:iNdEx+skippy]...)
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func skipGenerated(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowGenerated
			}
			if iNdEx >= l {
				return 0, io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		wireType := int(wire & 0x7)
		switch wireType {
		case 0:
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowGenerated
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				iNdEx++
				if dAtA[iNdEx-1] < 0x80 {
					break
				}
			}
			return iNdEx, nil
		case 1:
			iNdEx += 8
			return iNdEx, nil
		case 2:
			var length int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowGenerated
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				length |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			iNdEx += length
			if length < 0 {
				return 0, ErrInvalidLengthGenerated
			}
			return iNdEx, nil
		case 3:
			for {
				var innerWire uint64
				var start int = iNdEx
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return 0, ErrIntOverflowGenerated
					}
					if iNdEx >= l {
						return 0, io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					innerWire |= (uint64(b) & 0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				innerWireType := int(innerWire & 0x7)
				if innerWireType == 4 {
					break
				}
				next, err := skipGenerated(dAtA[start:])
				if err != nil {
					return 0, err
				}
				iNdEx = start + next
			}
			return iNdEx, nil
		case 4:
			return iNdEx, nil
		case 5:
			iNdEx += 4
			return iNdEx, nil
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
	}
	panic("unreachable")
}

var (
	ErrInvalidLengthGenerated = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowGenerated   = fmt.Errorf("proto: integer overflow")
)

func init() {
	proto.RegisterFile("k8s.io/api/imagepolicy/v1alpha1/generated.proto", fileDescriptorGenerated)
}

var fileDescriptorGenerated = []byte{
	// 446 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x91, 0xc1, 0x8e, 0xd3, 0x30,
	0x10, 0x86, 0x49, 0x0b, 0x4b, 0xeb, 0x1e, 0x58, 0x2c, 0x84, 0xa2, 0x0a, 0x95, 0x55, 0x4f, 0x7b,
	0xb2, 0xb7, 0x05, 0xa1, 0x85, 0x03, 0xd2, 0x02, 0x7b, 0x00, 0x81, 0x90, 0xcc, 0x09, 0x6e, 0x83,
	0x3b, 0x6a, 0x4d, 0x13, 0xdb, 0x8a, 0xdd, 0xac, 0xfa, 0x26, 0xbc, 0x05, 0xaf, 0xc1, 0x91, 0x47,
	0x40, 0xe5, 0xcc, 0x3b, 0x20, 0xbb, 0xed, 0x26, 0x34, 0x54, 0xa8, 0xb7, 0xcc, 0xc4, 0xdf, 0xe7,
	0x99, 0xdf, 0x84, 0xcf, 0xcf, 0x1d, 0x53, 0x86, 0x83, 0x55, 0x5c, 0xe5, 0x30, 0x45, 0x6b, 0x32,
	0x25, 0x97, 0xbc, 0x1c, 0x41, 0x66, 0x67, 0x30, 0xe2, 0x53, 0xd4, 0x58, 0x80, 0xc7, 0x09, 0xb3,
	0x85, 0xf1, 0x86, 0x3e, 0x5c, 0x03, 0x0c, 0xac, 0x62, 0x35, 0x80, 0x6d, 0x81, 0xfe, 0xe3, 0xca,
	0x98, 0x83, 0x9c, 0x29, 0x8d, 0xc5, 0x92, 0xdb, 0xf9, 0x34, 0x34, 0x1c, 0xcf, 0xd1, 0x03, 0x2f,
	0x1b, 0xda, 0x3e, 0xdf, 0x47, 0x15, 0x0b, 0xed, 0x55, 0x8e, 0x0d, 0xe0, 0xc9, 0xff, 0x00, 0x27,
	0x67, 0x98, 0x43, 0x83, 0x7b, 0xb4, 0x8f, 0x5b, 0x78, 0x95, 0x71, 0xa5, 0xbd, 0xf3, 0xc5, 0x2e,
	0x34, 0xfc, 0x9d, 0x90, 0xde, 0xeb, 0xb0, 0xac, 0xc0, 0x52, 0xe1, 0x15, 0x7d, 0x4b, 0x3a, 0x61,
	0x91, 0x09, 0x78, 0x48, 0x93, 0x93, 0xe4, 0xb4, 0x37, 0x3e, 0x63, 0x55, 0x2e, 0xd7, 0x5e, 0x66,
	0xe7, 0xd3, 0xd0, 0x70, 0x2c, 0x9c, 0x66, 0xe5, 0x88, 0xbd, 0xff, 0xfc, 0x05, 0xa5, 0x7f, 0x87,
	0x1e, 0xc4, 0xb5, 0x81, 0xbe, 0x22, 0x37, 0x9d, 0x45, 0x99, 0xb6, 0x1a, 0xa6, 0x7f, 0x26, 0xcc,
	0x6a, 0x93, 0x7c, 0xb0, 0x28, 0x45, 0xa4, 0xe9, 0x1b, 0x72, 0xe4, 0x3c, 0xf8, 0x85, 0x4b, 0xdb,
	0xd1, 0x33, 0x3e, 0xc8, 0x13, 0x49, 0xb1, 0x31, 0x0c, 0xcf, 0x48, 0x5a, 0xfb, 0xf9, 0xd2, 0x68,
	0x0f, 0x61, 0xa1, 0x70, 0x1b, 0xbd, 0x47, 0x6e, 0x45, 0x5b, 0x5c, 0xbc, 0x2b, 0xd6, 0xc5, 0xf0,
	0x5b, 0x8b, 0xdc, 0xd9, 0x99, 0x8b, 0x7e, 0x24, 0x44, 0x6e, 0x51, 0x97, 0x26, 0x27, 0xed, 0xd3,
	0xde, 0xf8, 0xe9, 0x21, 0x53, 0xfd, 0x75, 0xb1, 0xa8, 0xc9, 0xa8, 0x24, 0x3d, 0xd0, 0xda, 0x78,
	0xf0, 0xca, 0x68, 0x97, 0xb6, 0xa2, 0xfb, 0xe2, 0xd0, 0xe4, 0xd8, 0x45, 0xe5, 0xb8, 0xd4, 0xbe,
	0x58, 0x8a, 0xba, 0x95, 0x3e, 0x20, 0x5d, 0x0d, 0x39, 0x3a, 0x0b, 0x12, 0x63, 0xa8, 0x5d, 0x51,
	0x35, 0xfa, 0xcf, 0xc9, 0xf1, 0x2e, 0x4e, 0x8f, 0x49, 0x7b, 0x8e, 0xcb, 0x4d, 0x32, 0xe1, 0x33,
	0xa4, 0x55, 0x42, 0xb6, 0xc0, 0xf8, 0xb8, 0x5d, 0xb1, 0x2e, 0x9e, 0xb5, 0xce, 0x93, 0xe1, 0x25,
	0xb9, 0xdb, 0x78, 0x00, 0x9a, 0x92, 0xdb, 0x90, 0x65, 0xe6, 0x0a, 0x27, 0x51, 0xd2, 0x11, 0xdb,
	0x92, 0xde, 0x27, 0x47, 0x05, 0x82, 0x33, 0x7a, 0x63, 0xda, 0x54, 0x2f, 0xfa, 0xdf, 0x57, 0x83,
	0xe4, 0xc7, 0x6a, 0x90, 0xfc, 0x5c, 0x0d, 0x92, 0xaf, 0xbf, 0x06, 0x37, 0x3e, 0x75, 0xb6, 0xeb,
	0xfe, 0x09, 0x00, 0x00, 0xff, 0xff, 0x40, 0x56, 0x67, 0xe0, 0xdd, 0x03, 0x00, 0x00,
}
