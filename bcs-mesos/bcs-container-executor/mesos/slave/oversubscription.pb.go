/*
 * Tencent is pleased to support the open source community by making Blueking Container Service available.
 * Copyright (C) 2019 THL A29 Limited, a Tencent company. All rights reserved.
 * Licensed under the MIT License (the "License"); you may not use this file except
 * in compliance with the License. You may obtain a copy of the License at
 * http://opensource.org/licenses/MIT
 * Unless required by applicable law or agreed to in writing, software distributed under
 * the License is distributed on an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND,
 * either express or implied. See the License for the specific language governing permissions and
 * limitations under the License.
 *
 */

// Code generated by protoc-gen-go.
// source: mesos/slave/oversubscription.proto
// DO NOT EDIT!

package slave

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import mesos "mesos"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// NOTE: In the future we can define more actions like
// freeze and resize.
type QoSCorrection_Type int32

const (
	QoSCorrection_KILL QoSCorrection_Type = 1
)

var QoSCorrection_Type_name = map[int32]string{
	1: "KILL",
}
var QoSCorrection_Type_value = map[string]int32{
	"KILL": 1,
}

func (x QoSCorrection_Type) Enum() *QoSCorrection_Type {
	p := new(QoSCorrection_Type)
	*p = x
	return p
}
func (x QoSCorrection_Type) String() string {
	return proto.EnumName(QoSCorrection_Type_name, int32(x))
}
func (x *QoSCorrection_Type) UnmarshalJSON(data []byte) error {
	value, err := proto.UnmarshalJSONEnum(QoSCorrection_Type_value, data, "QoSCorrection_Type")
	if err != nil {
		return err
	}
	*x = QoSCorrection_Type(value)
	return nil
}
func (QoSCorrection_Type) EnumDescriptor() ([]byte, []int) { return fileDescriptor1, []int{0, 0} }

// *
// The QoS Controller informs the slave that particular corrective
// actions need to be made. Each corrective action contains
// information about executor or task and the type of action to
// perform.
type QoSCorrection struct {
	Type             *QoSCorrection_Type `protobuf:"varint,1,req,name=type,enum=mesos.slave.QoSCorrection_Type" json:"type,omitempty"`
	Kill             *QoSCorrection_Kill `protobuf:"bytes,2,opt,name=kill" json:"kill,omitempty"`
	XXX_unrecognized []byte              `json:"-"`
}

func (m *QoSCorrection) Reset()                    { *m = QoSCorrection{} }
func (m *QoSCorrection) String() string            { return proto.CompactTextString(m) }
func (*QoSCorrection) ProtoMessage()               {}
func (*QoSCorrection) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{0} }

func (m *QoSCorrection) GetType() QoSCorrection_Type {
	if m != nil && m.Type != nil {
		return *m.Type
	}
	return QoSCorrection_KILL
}

func (m *QoSCorrection) GetKill() *QoSCorrection_Kill {
	if m != nil {
		return m.Kill
	}
	return nil
}

// Kill action which will be performed on an executor.
// NOTE: Framework ID and either executor id or executor id and container_id
// must be set to kill a running container. In the latter case, the caller
// can be assured that the kill request doesn't target a new executor reusing
// an old executor id.
// NOTE: In the future we may also kill individual tasks by
// specifying an optional 'task_id'.
type QoSCorrection_Kill struct {
	FrameworkId      *mesos.FrameworkID `protobuf:"bytes,1,opt,name=framework_id,json=frameworkId" json:"framework_id,omitempty"`
	ExecutorId       *mesos.ExecutorID  `protobuf:"bytes,2,opt,name=executor_id,json=executorId" json:"executor_id,omitempty"`
	ContainerId      *mesos.ContainerID `protobuf:"bytes,3,opt,name=container_id,json=containerId" json:"container_id,omitempty"`
	XXX_unrecognized []byte             `json:"-"`
}

func (m *QoSCorrection_Kill) Reset()                    { *m = QoSCorrection_Kill{} }
func (m *QoSCorrection_Kill) String() string            { return proto.CompactTextString(m) }
func (*QoSCorrection_Kill) ProtoMessage()               {}
func (*QoSCorrection_Kill) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{0, 0} }

func (m *QoSCorrection_Kill) GetFrameworkId() *mesos.FrameworkID {
	if m != nil {
		return m.FrameworkId
	}
	return nil
}

func (m *QoSCorrection_Kill) GetExecutorId() *mesos.ExecutorID {
	if m != nil {
		return m.ExecutorId
	}
	return nil
}

func (m *QoSCorrection_Kill) GetContainerId() *mesos.ContainerID {
	if m != nil {
		return m.ContainerId
	}
	return nil
}

func init() {
	proto.RegisterType((*QoSCorrection)(nil), "mesos.slave.QoSCorrection")
	proto.RegisterType((*QoSCorrection_Kill)(nil), "mesos.slave.QoSCorrection.Kill")
	proto.RegisterEnum("mesos.slave.QoSCorrection_Type", QoSCorrection_Type_name, QoSCorrection_Type_value)
}

func init() { proto.RegisterFile("mesos/slave/oversubscription.proto", fileDescriptor1) }

var fileDescriptor1 = []byte{
	// 269 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x7c, 0x90, 0xb1, 0x4e, 0xf3, 0x30,
	0x14, 0x85, 0x95, 0xfc, 0xd1, 0x2f, 0x74, 0x03, 0x28, 0xf5, 0x14, 0x65, 0xa1, 0xca, 0xd4, 0xc9,
	0x95, 0x8a, 0x78, 0x81, 0xb6, 0x20, 0x59, 0xed, 0x00, 0x81, 0x1d, 0x85, 0xe4, 0x02, 0x56, 0xd3,
	0x5c, 0xcb, 0x76, 0x0b, 0x7d, 0x23, 0x9e, 0x80, 0xe7, 0x43, 0x76, 0x92, 0x8a, 0x2e, 0x8c, 0xf6,
	0xf9, 0xbe, 0x73, 0x64, 0x43, 0xbe, 0x45, 0x43, 0x66, 0x6a, 0x9a, 0x72, 0x8f, 0x53, 0xda, 0xa3,
	0x36, 0xbb, 0x17, 0x53, 0x69, 0xa9, 0xac, 0xa4, 0x96, 0x2b, 0x4d, 0x96, 0x58, 0xec, 0x19, 0xee,
	0x99, 0x6c, 0xd4, 0x09, 0xdd, 0x95, 0xcf, 0xf3, 0xef, 0x10, 0x2e, 0x1e, 0xe8, 0x71, 0x41, 0x5a,
	0x63, 0xe5, 0x3c, 0x76, 0x0d, 0x91, 0x3d, 0x28, 0x4c, 0x83, 0x71, 0x38, 0xb9, 0x9c, 0x5d, 0xf1,
	0x5f, 0x05, 0xfc, 0x84, 0xe4, 0x4f, 0x07, 0x85, 0x85, 0x87, 0x9d, 0xb4, 0x91, 0x4d, 0x93, 0x86,
	0xe3, 0x60, 0x12, 0xff, 0x29, 0xad, 0x64, 0xd3, 0x14, 0x1e, 0xce, 0xbe, 0x02, 0x88, 0xdc, 0x91,
	0xdd, 0xc0, 0xf9, 0xab, 0x2e, 0xb7, 0xf8, 0x41, 0x7a, 0xf3, 0x2c, 0xeb, 0x34, 0xf0, 0x2d, 0xac,
	0x6f, 0xb9, 0x1b, 0x22, 0xb1, 0x2c, 0xe2, 0x23, 0x27, 0x6a, 0x36, 0x83, 0x18, 0x3f, 0xb1, 0xda,
	0x59, 0xd2, 0xce, 0xea, 0xb6, 0x47, 0xbd, 0x75, 0xdb, 0x27, 0x62, 0x59, 0xc0, 0x40, 0x89, 0xda,
	0x4d, 0x55, 0xd4, 0xda, 0x52, 0xb6, 0xe8, 0xa5, 0x7f, 0x27, 0x53, 0x8b, 0x21, 0x72, 0x53, 0x47,
	0x4e, 0xd4, 0x79, 0x02, 0x91, 0x7b, 0x2d, 0x3b, 0x83, 0x68, 0x25, 0xd6, 0xeb, 0x24, 0x98, 0x67,
	0x90, 0x90, 0x7e, 0xe3, 0xa5, 0x2a, 0xab, 0x77, 0xec, 0xf4, 0xf9, 0xff, 0x7b, 0xf7, 0xa7, 0xe6,
	0x27, 0x00, 0x00, 0xff, 0xff, 0x30, 0x16, 0x35, 0xac, 0x99, 0x01, 0x00, 0x00,
}
