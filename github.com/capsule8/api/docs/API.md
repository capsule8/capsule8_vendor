# Protocol Documentation
<a name="top"/>

## Table of Contents

- [types.proto](#types.proto)
    - [IPv4Address](#capsule8.api.v0.IPv4Address)
    - [IPv4AddressAndPort](#capsule8.api.v0.IPv4AddressAndPort)
    - [IPv6Address](#capsule8.api.v0.IPv6Address)
    - [IPv6AddressAndPort](#capsule8.api.v0.IPv6AddressAndPort)
    - [NetworkAddress](#capsule8.api.v0.NetworkAddress)
  
    - [NetworkAddressFamily](#capsule8.api.v0.NetworkAddressFamily)
  
  
  

- [event.proto](#event.proto)
    - [ChargenEvent](#capsule8.api.v0.ChargenEvent)
    - [ContainerEvent](#capsule8.api.v0.ContainerEvent)
    - [Event](#capsule8.api.v0.Event)
    - [FileEvent](#capsule8.api.v0.FileEvent)
    - [KernelFunctionCallEvent](#capsule8.api.v0.KernelFunctionCallEvent)
    - [KernelFunctionCallEvent.ArgumentsEntry](#capsule8.api.v0.KernelFunctionCallEvent.ArgumentsEntry)
    - [KernelFunctionCallEvent.FieldValue](#capsule8.api.v0.KernelFunctionCallEvent.FieldValue)
    - [NetworkEvent](#capsule8.api.v0.NetworkEvent)
    - [Process](#capsule8.api.v0.Process)
    - [ProcessEvent](#capsule8.api.v0.ProcessEvent)
    - [SyscallEvent](#capsule8.api.v0.SyscallEvent)
    - [TickerEvent](#capsule8.api.v0.TickerEvent)
  
    - [ContainerEventType](#capsule8.api.v0.ContainerEventType)
    - [FileEventType](#capsule8.api.v0.FileEventType)
    - [KernelFunctionCallEvent.FieldType](#capsule8.api.v0.KernelFunctionCallEvent.FieldType)
    - [KernelFunctionCallEventType](#capsule8.api.v0.KernelFunctionCallEventType)
    - [NetworkEventType](#capsule8.api.v0.NetworkEventType)
    - [ProcessEventType](#capsule8.api.v0.ProcessEventType)
    - [SyscallEventType](#capsule8.api.v0.SyscallEventType)
  
  
  

- [subscription.proto](#subscription.proto)
    - [ChargenEventFilter](#capsule8.api.v0.ChargenEventFilter)
    - [ContainerEventFilter](#capsule8.api.v0.ContainerEventFilter)
    - [ContainerFilter](#capsule8.api.v0.ContainerFilter)
    - [EventFilter](#capsule8.api.v0.EventFilter)
    - [FileEventFilter](#capsule8.api.v0.FileEventFilter)
    - [FilterExpression](#capsule8.api.v0.FilterExpression)
    - [FilterPredicate](#capsule8.api.v0.FilterPredicate)
    - [KernelFunctionCallFilter](#capsule8.api.v0.KernelFunctionCallFilter)
    - [KernelFunctionCallFilter.ArgumentsEntry](#capsule8.api.v0.KernelFunctionCallFilter.ArgumentsEntry)
    - [LimitModifier](#capsule8.api.v0.LimitModifier)
    - [Modifier](#capsule8.api.v0.Modifier)
    - [NetworkEventFilter](#capsule8.api.v0.NetworkEventFilter)
    - [ProcessEventFilter](#capsule8.api.v0.ProcessEventFilter)
    - [Subscription](#capsule8.api.v0.Subscription)
    - [SyscallEventFilter](#capsule8.api.v0.SyscallEventFilter)
    - [ThrottleModifier](#capsule8.api.v0.ThrottleModifier)
    - [TickerEventFilter](#capsule8.api.v0.TickerEventFilter)
  
    - [ContainerEventView](#capsule8.api.v0.ContainerEventView)
    - [FilterExpression.FilterExpressionType](#capsule8.api.v0.FilterExpression.FilterExpressionType)
    - [FilterPredicate.PredicateType](#capsule8.api.v0.FilterPredicate.PredicateType)
    - [FilterPredicate.ValueType](#capsule8.api.v0.FilterPredicate.ValueType)
    - [ThrottleModifier.IntervalType](#capsule8.api.v0.ThrottleModifier.IntervalType)
  
  
  

- [telemetry_service.proto](#telemetry_service.proto)
    - [GetEventsRequest](#capsule8.api.v0.GetEventsRequest)
    - [GetEventsResponse](#capsule8.api.v0.GetEventsResponse)
    - [TelemetryEvent](#capsule8.api.v0.TelemetryEvent)
  
  
  
    - [TelemetryService](#capsule8.api.v0.TelemetryService)
  

- [Scalar Value Types](#scalar-value-types)



<a name="types.proto"/>
<p align="right"><a href="#top">Top</a></p>

## types.proto



<a name="capsule8.api.v0.IPv4Address"/>

### IPv4Address
An IPv4 address


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| address | [fixed32](#fixed32) |  | The IPv4 address is network byte order (big endian) |






<a name="capsule8.api.v0.IPv4AddressAndPort"/>

### IPv4AddressAndPort
An IPv4 address and port


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| address | [IPv4Address](#capsule8.api.v0.IPv4Address) |  | The IPv4 address |
| port | [uint32](#uint32) |  | The port |






<a name="capsule8.api.v0.IPv6Address"/>

### IPv6Address
An IPv6 address


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| high | [fixed64](#fixed64) |  | The high-order bytes of the IPv6 address |
| low | [fixed64](#fixed64) |  | The low-order bytes of the IPv6 address |






<a name="capsule8.api.v0.IPv6AddressAndPort"/>

### IPv6AddressAndPort
An IPv6 address and port


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| address | [IPv6Address](#capsule8.api.v0.IPv6Address) |  | The IPv6 address |
| port | [uint32](#uint32) |  | The port |






<a name="capsule8.api.v0.NetworkAddress"/>

### NetworkAddress
A network address


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| family | [NetworkAddressFamily](#capsule8.api.v0.NetworkAddressFamily) |  | The address family that specifies which address format is in use |
| ipv4_address | [IPv4AddressAndPort](#capsule8.api.v0.IPv4AddressAndPort) |  | Used when family is NETWORK_ADDRESS_INET |
| ipv6_address | [IPv6AddressAndPort](#capsule8.api.v0.IPv6AddressAndPort) |  | Used when family is NETWORK_ADDRESS_INET6 |
| local_address | [string](#string) |  | Used when family is NETWORK_ADDRESS_LOCAL |





 


<a name="capsule8.api.v0.NetworkAddressFamily"/>

### NetworkAddressFamily
Supported network address families

| Name | Number | Description |
| ---- | ------ | ----------- |
| NETWORK_ADDRESS_FAMILY_UNKNOWN | 0 | The network address family is unknown |
| NETWORK_ADDRESS_FAMILY_INET | 1 | AF_INET; IPv4 address formats |
| NETWORK_ADDRESS_FAMILY_INET6 | 2 | AF_INET6; IPv6 address formats |
| NETWORK_ADDRESS_FAMILY_LOCAL | 3 | AF_LOCAL / AF_UNIX; local filesystem address formats |


 

 

 



<a name="event.proto"/>
<p align="right"><a href="#top">Top</a></p>

## event.proto



<a name="capsule8.api.v0.ChargenEvent"/>

### ChargenEvent



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| index | [uint64](#uint64) |  | Index of the first character in this Event in relation to all of the characters that have been generated in this stream. |
| characters | [string](#string) |  | The next one or more characters in the autogenerated stream |






<a name="capsule8.api.v0.ContainerEvent"/>

### ContainerEvent
ContainerEvent describes a Docker container or Rkt App lifecycle event


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| type | [ContainerEventType](#capsule8.api.v0.ContainerEventType) |  |  |
| name | [string](#string) |  |  |
| image_id | [string](#string) |  | Unique identifier of the container image |
| image_name | [string](#string) |  | Name of the container image (i.e. &#34;busybox&#34; or &#34;gcr.io/google_containers/nginx-ingress-controller&#34;) |
| host_pid | [sint32](#sint32) |  | Host process identifier of the container&#39;s init process. |
| exit_code | [sint32](#sint32) |  | Optional, only included on CONTAINER_EVENT_TYPE_EXIT events |
| exit_status | [uint32](#uint32) |  | The exit status will typically one of the values defined in stdlib.h like EXIT_SUCCESS, EXIT_FAILURE, or EXIT_USAGE. |
| exit_signal | [uint32](#uint32) |  | If non-zero, this is the signal number that the process was terminated with. |
| exit_core_dumped | [bool](#bool) |  | If true, indicates that the process dumped a core when it terminated. |
| docker_config_json | [string](#string) |  | Docker container configuration file |
| oci_config_json | [string](#string) |  | OCI container configuration file |






<a name="capsule8.api.v0.Event"/>

### Event
An event observed by the Sensor.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| id | [string](#string) |  | Unique identifier for the event |
| process_id | [string](#string) |  | Unique process identifier associated with the event to differentiate reused values of the pid below. |
| process_pid | [int32](#int32) |  | Unix pid of the process associated with the event |
| container_id | [string](#string) |  | Container identifier associated with the event |
| sensor_id | [string](#string) |  | Sensor identifier of the sensor instance that observed the event |
| sensor_sequence_number | [uint64](#uint64) |  | Sequence number from some unspecified starting point unique to the Sensor. Provides a strict linear ordering of events with the same sensor_id where no two events can have the same sequence number. If it is present, it must be greater than zero. A zero value indicates that there is no sequence number associated with the event. |
| sensor_monotime_nanos | [int64](#int64) |  | Monotonic nanosecond timestamp from some unspecified starting point unique to the Sensor. Can only be used to calculate time intervals between events with the same sensor_id. |
| process_lineage | [Process](#capsule8.api.v0.Process) | repeated | Process Lineage contains one process context for each process in the hierarchy, starting with the current process, up to the root of the process namespace. |
| container_name | [string](#string) |  | Name of container associated with the event |
| image_id | [string](#string) |  | Unique identifier of the container image |
| image_name | [string](#string) |  | Name of the container image (i.e. &#34;busybox&#34; or &#34;gcr.io/google_containers/nginx-ingress-controller&#34;) |
| syscall | [SyscallEvent](#capsule8.api.v0.SyscallEvent) |  |  |
| process | [ProcessEvent](#capsule8.api.v0.ProcessEvent) |  |  |
| file | [FileEvent](#capsule8.api.v0.FileEvent) |  |  |
| kernel_call | [KernelFunctionCallEvent](#capsule8.api.v0.KernelFunctionCallEvent) |  |  |
| network | [NetworkEvent](#capsule8.api.v0.NetworkEvent) |  |  |
| container | [ContainerEvent](#capsule8.api.v0.ContainerEvent) |  |  |
| chargen | [ChargenEvent](#capsule8.api.v0.ChargenEvent) |  | Debugging events (&gt;= 100) |
| ticker | [TickerEvent](#capsule8.api.v0.TickerEvent) |  |  |
| process_tid | [int32](#int32) |  | Unix tid of the process thread associated with event |
| cpu | [int32](#int32) |  | CPU on which the event occurred |






<a name="capsule8.api.v0.FileEvent"/>

### FileEvent
FileEvent describes an event that occurred related to file operations
occurring as detected by the Sensor.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| type | [FileEventType](#capsule8.api.v0.FileEventType) |  | The type of event described by this FileEvent message |
| filename | [string](#string) |  | Present when the event is a file open event. This is the filename of the file being opened. |
| open_flags | [sint32](#sint32) |  | Present when the event is a file open event. This is the set of flags with which the file was opened (e.g., O_RDONLY, O_NONBLOCK, etc.). |
| open_mode | [sint32](#sint32) |  | Present when the event is a file open event. This is the set of file permissions used in a creat(2) system call. |






<a name="capsule8.api.v0.KernelFunctionCallEvent"/>

### KernelFunctionCallEvent
KernelFunctionCallEvent describes an event that occurred related to kernel
functions being entered or exited.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| arguments | [KernelFunctionCallEvent.ArgumentsEntry](#capsule8.api.v0.KernelFunctionCallEvent.ArgumentsEntry) | repeated | Label repeated w/ a `mapEntry` option set to `true`. This is a map of argument names and values. The keys are strings that are the names of the arguments, and the values are the actual values for each field. |






<a name="capsule8.api.v0.KernelFunctionCallEvent.ArgumentsEntry"/>

### KernelFunctionCallEvent.ArgumentsEntry



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| key | [string](#string) |  |  |
| value | [KernelFunctionCallEvent.FieldValue](#capsule8.api.v0.KernelFunctionCallEvent.FieldValue) |  |  |






<a name="capsule8.api.v0.KernelFunctionCallEvent.FieldValue"/>

### KernelFunctionCallEvent.FieldValue
The representation of a field value, which is composed of type
information and the value itself.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| field_type | [KernelFunctionCallEvent.FieldType](#capsule8.api.v0.KernelFunctionCallEvent.FieldType) |  | The type represented by this field value. |
| bytes_value | [bytes](#bytes) |  | An array of bytes |
| string_value | [string](#string) |  | A string |
| signed_value | [sint64](#sint64) |  | A signed value (8-bit, 16-bit, 32-bit, or 64-bit) |
| unsigned_value | [uint64](#uint64) |  | An unsigned value (8-bit, 16--bit, 32-bit, or 64-bit) |






<a name="capsule8.api.v0.NetworkEvent"/>

### NetworkEvent
NetworkEvent describes an event that occurred related to network activity
occurring as detected by the Sensor.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| type | [NetworkEventType](#capsule8.api.v0.NetworkEventType) |  | The type of event described by this NetworkEvent message. |
| sockfd | [uint64](#uint64) |  | Present when the event describes a network event that is an attempt to perform a network related action. This is the socket descriptor used to perform the action. |
| address | [NetworkAddress](#capsule8.api.v0.NetworkAddress) |  | Present when the event describes a network event that is an attempt to perform a network related action that includes an address. This is that address. |
| result | [sint64](#sint64) |  | Present when the event describes a network event that is the result of an attempted network related action. This is the return code from the system call. |
| backlog | [uint64](#uint64) |  | Present only when the event describes a listen attempt. This is the value of the backlog argument passed to listen(2). |






<a name="capsule8.api.v0.Process"/>

### Process



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| pid | [sint32](#sint32) |  |  |
| command | [string](#string) |  |  |






<a name="capsule8.api.v0.ProcessEvent"/>

### ProcessEvent
ProcessEvent describes an event that occurred related to processes starting
and exiting as detected by the Sensor.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| type | [ProcessEventType](#capsule8.api.v0.ProcessEventType) |  | The type of event described by this ProcessEvent message |
| fork_child_pid | [sint32](#sint32) |  | Present when the event is a fork event. This is the PID of the new child process. |
| fork_child_id | [string](#string) |  | Present when the event is a fork event. This is the Sensor&#39;s process ID of the new child process. |
| exec_filename | [string](#string) |  | Present when the event is an exec event. This is the filename of the executable that was executed. |
| exec_command_line | [string](#string) | repeated | Present when the event is an exec event. Repeated for each argument passed to the executable on the command-line. |
| exit_code | [sint32](#sint32) |  | Present when the event is an exit event. This is the exit code that the process exited with. |
| exit_status | [uint32](#uint32) |  | Present when the event is an exit event. This will typically be one9 of the values defined in stdlib.h like EXIT_SUCCESS, EXIT_FAILURE, or EXIT_USAGE. |
| exit_signal | [uint32](#uint32) |  | Present when the event is an exit event. If non-zero, this is the signal number that the process was terminated with. |
| exit_core_dumped | [bool](#bool) |  | Present when the event is an exit event. If true, indicates that the process dumped a core when it terminated. |






<a name="capsule8.api.v0.SyscallEvent"/>

### SyscallEvent
SyscallEvent describes an event that occurred related to system calls being
made or returning as detected by the Sensor.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| type | [SyscallEventType](#capsule8.api.v0.SyscallEventType) |  | The type of event described by this SyscallEvent message |
| id | [int64](#int64) |  | The syscall number for either enter or exit events. |
| arg0 | [uint64](#uint64) |  | Present when the event is an enter event. This is the first argument passed to the system call. |
| arg1 | [uint64](#uint64) |  | Present when the event is an enter event. This is the second argument passed to the system call. |
| arg2 | [uint64](#uint64) |  | Present when the event is an enter event. This is the third argument passed to the system call. |
| arg3 | [uint64](#uint64) |  | Present when the event is an enter event. This is the fourth argument passed to the system call. |
| arg4 | [uint64](#uint64) |  | Present when the event is an enter event. This is the fifth argument passed to the system call. |
| arg5 | [uint64](#uint64) |  | Present when the event is an enter event. This is the sixth argument passed to the system call. |
| ret | [int64](#int64) |  | Present when the event is an exit event. This is the value that was returned from the system call. |






<a name="capsule8.api.v0.TickerEvent"/>

### TickerEvent



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| seconds | [int64](#int64) |  | The number of seconds elapsed since January 1, 1970 UTC.

https://golang.org/pkg/time/#Time.Unix |
| nanoseconds | [int64](#int64) |  | The number of nanoseconds elapsed since January 1, 1970 UTC

https://golang.org/pkg/time/#Time.UnixNano |





 


<a name="capsule8.api.v0.ContainerEventType"/>

### ContainerEventType


| Name | Number | Description |
| ---- | ------ | ----------- |
| CONTAINER_EVENT_TYPE_UNKNOWN | 0 |  |
| CONTAINER_EVENT_TYPE_CREATED | 1 |  |
| CONTAINER_EVENT_TYPE_RUNNING | 2 |  |
| CONTAINER_EVENT_TYPE_EXITED | 3 |  |
| CONTAINER_EVENT_TYPE_DESTROYED | 4 |  |



<a name="capsule8.api.v0.FileEventType"/>

### FileEventType
Possible FileEvent types

| Name | Number | Description |
| ---- | ------ | ----------- |
| FILE_EVENT_TYPE_UNKNOWN | 0 | The type of event is unknown |
| FILE_EVENT_TYPE_OPEN | 1 | The event is a file open event |



<a name="capsule8.api.v0.KernelFunctionCallEvent.FieldType"/>

### KernelFunctionCallEvent.FieldType
Possible field types

| Name | Number | Description |
| ---- | ------ | ----------- |
| UNKNOWN | 0 | The field type is unknown |
| BYTES | 1 | The field type is an array of bytes |
| STRING | 2 | The field type is a string |
| SINT8 | 3 | The field type is a signed 8-bit integer |
| SINT16 | 4 | The field type is a signed 16-bit integer |
| SINT32 | 5 | The field type is a signed 32-bit integer |
| SINT64 | 6 | The field type is a signed 64-bit integer |
| UINT8 | 7 | The field type is an unsigned 8-bit integer |
| UINT16 | 8 | The field type is an unsigned 16-bit integer |
| UINT32 | 9 | The field type is an unsigned 32-bit integer |
| UINT64 | 10 | The field type is an unsigned 64-bit integer |



<a name="capsule8.api.v0.KernelFunctionCallEventType"/>

### KernelFunctionCallEventType
Possible KernelFunctionCallEvent types

| Name | Number | Description |
| ---- | ------ | ----------- |
| KERNEL_FUNCTION_CALL_EVENT_TYPE_UNKNOWN | 0 | The type of event is unknown |
| KERNEL_FUNCTION_CALL_EVENT_TYPE_ENTER | 1 | The event is a kernel function being entered. |
| KERNEL_FUNCTION_CALL_EVENT_TYPE_EXIT | 2 | The event is a kernel function being exited. |



<a name="capsule8.api.v0.NetworkEventType"/>

### NetworkEventType
Possible network event types

| Name | Number | Description |
| ---- | ------ | ----------- |
| NETWORK_EVENT_TYPE_UNKNOWN | 0 | The type of event is unknown |
| NETWORK_EVENT_TYPE_CONNECT_ATTEMPT | 1 | The event is an attempt to connect to an address |
| NETWORK_EVENT_TYPE_CONNECT_RESULT | 2 | The event is the result of an attempt to connect to an address |
| NETWORK_EVENT_TYPE_BIND_ATTEMPT | 3 | The event is an attempt to bind to a local address |
| NETWORK_EVENT_TYPE_BIND_RESULT | 4 | The event is the result of an attempt to bind to a local address |
| NETWORK_EVENT_TYPE_LISTEN_ATTEMPT | 5 | The event is an attempt to listen for connections |
| NETWORK_EVENT_TYPE_LISTEN_RESULT | 6 | The event is the result of an attempt to listen for connections |
| NETWORK_EVENT_TYPE_ACCEPT_ATTEMPT | 7 | The event is an attempt to accept an incoming connection |
| NETWORK_EVENT_TYPE_ACCEPT_RESULT | 8 | The event is the result of an attempt to accept an incoming connection |
| NETWORK_EVENT_TYPE_SENDTO_ATTEMPT | 9 | The event is an attempt to send data to a specific address |
| NETWORK_EVENT_TYPE_SENDTO_RESULT | 10 | The event is the result of an attempt to send data to a specific address |
| NETWORK_EVENT_TYPE_RECVFROM_ATTEMPT | 11 | The event is an attempt to receive data from a specific address |
| NETWORK_EVENT_TYPE_RECVFROM_RESULT | 12 | The event is the result of an attempt to receive data from a specific address |



<a name="capsule8.api.v0.ProcessEventType"/>

### ProcessEventType
Possible ProcessEvent types

| Name | Number | Description |
| ---- | ------ | ----------- |
| PROCESS_EVENT_TYPE_UNKNOWN | 0 | The type of event is unknown |
| PROCESS_EVENT_TYPE_FORK | 1 | The event is a process fork event |
| PROCESS_EVENT_TYPE_EXEC | 2 | The event is a process exec event |
| PROCESS_EVENT_TYPE_EXIT | 3 | The event is a process exit event |



<a name="capsule8.api.v0.SyscallEventType"/>

### SyscallEventType
Possible SyscallEvent types

| Name | Number | Description |
| ---- | ------ | ----------- |
| SYSCALL_EVENT_TYPE_UNKNOWN | 0 | The type of event is unknown |
| SYSCALL_EVENT_TYPE_ENTER | 1 | The event is a syscall enter event |
| SYSCALL_EVENT_TYPE_EXIT | 2 | The event is a syscall exit event |


 

 

 



<a name="subscription.proto"/>
<p align="right"><a href="#top">Top</a></p>

## subscription.proto



<a name="capsule8.api.v0.ChargenEventFilter"/>

### ChargenEventFilter
The ChargenEventFilter configures a character stream generator and
includes events from it in the Subscription.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| length | [uint64](#uint64) |  | Required; the length of character sequence strings to generate |






<a name="capsule8.api.v0.ContainerEventFilter"/>

### ContainerEventFilter
The ContainerEventFilter specifies which container lifecycle events
to include in the Subscription. In order to restrict them to
specific containers, use the ContainerFilter.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| type | [ContainerEventType](#capsule8.api.v0.ContainerEventType) |  | Required, specify the particular type of event type to match |
| view | [ContainerEventView](#capsule8.api.v0.ContainerEventView) |  | Optional, specifies how much detail to include in container events |






<a name="capsule8.api.v0.ContainerFilter"/>

### ContainerFilter
The ContainerFilter restricts events in the Subscription to the
running containers indicated. All of the fields in this message are
effectively &#34;ORed&#34; together to create the list of containers to
monitor for the subscription.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| ids | [string](#string) | repeated | Zero or more container IDs (e.g. 254dd98a7bf1581560ddace9f98b7933bfb3c2f5fc0504ec1b8dcc9614bc7062) |
| names | [string](#string) | repeated | Zero or more container names (e.g. /ecstatic_darwin) |
| image_ids | [string](#string) | repeated | Zero or more container image IDs (e.g. d462265d362c919b7dd37f8ba80caa822d13704695f47c8fc42a1c2266ecd164) |
| image_names | [string](#string) | repeated | Container image name (shell-style globs are supported). May be of the form &#34;busybox&#34;, &#34;foo/bar&#34; or &#34;sha256:d462265d362c919b7dd37f8ba80caa822d13704695f47c8fc42a1c2266ecd164&#34; |






<a name="capsule8.api.v0.EventFilter"/>

### EventFilter
The EventFilter specifies events to include. All of the specified
fields are effectively &#34;ORed&#34; together to create the list of events
included in the Subscription.



Kernel-level events


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| syscall_events | [SyscallEventFilter](#capsule8.api.v0.SyscallEventFilter) | repeated | Zero or more filters specifying which system calls to include |
| process_events | [ProcessEventFilter](#capsule8.api.v0.ProcessEventFilter) | repeated | Zero or more filters specifying which process events to include |
| file_events | [FileEventFilter](#capsule8.api.v0.FileEventFilter) | repeated | Zero or more filters specifying which file events to include |
| kernel_events | [KernelFunctionCallFilter](#capsule8.api.v0.KernelFunctionCallFilter) | repeated | Zero or more kernel functional calls to include |
| network_events | [NetworkEventFilter](#capsule8.api.v0.NetworkEventFilter) | repeated | Zero or more network events to include |
| container_events | [ContainerEventFilter](#capsule8.api.v0.ContainerEventFilter) | repeated | Zero or more container events to include |
| chargen_events | [ChargenEventFilter](#capsule8.api.v0.ChargenEventFilter) | repeated | Zero or more character generators to configure and return events from (for debugging) |
| ticker_events | [TickerEventFilter](#capsule8.api.v0.TickerEventFilter) | repeated | Zero or more ticker generators to configure and return events from (for debugging) |






<a name="capsule8.api.v0.FileEventFilter"/>

### FileEventFilter
The FileEventFilter specifies which file events to include in the
Subscription. The specified fields are effectively &#34;ANDed&#34; to
specify a matching event.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| type | [FileEventType](#capsule8.api.v0.FileEventType) |  | Required; the file event type to match |
| filename | [.google.protobuf.StringValue](#capsule8.api.v0..google.protobuf.StringValue) |  | Optional; require exact match on the filename being acted upon |
| filename_pattern | [.google.protobuf.StringValue](#capsule8.api.v0..google.protobuf.StringValue) |  | Optional; require pattern match on the filename being acted upon |
| open_flags_mask | [.google.protobuf.Int32Value](#capsule8.api.v0..google.protobuf.Int32Value) |  | Optional; for file open events, require a match of the bits set for the open(2) flags argument |
| create_mode_mask | [.google.protobuf.Int32Value](#capsule8.api.v0..google.protobuf.Int32Value) |  | Optional; for file open events, require a match of the bits set for the open(2) or creat(2) mode argument |






<a name="capsule8.api.v0.FilterExpression"/>

### FilterExpression
A tree-based representation of an expression to be used to apply filters to
events.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| type | [FilterExpression.FilterExpressionType](#capsule8.api.v0.FilterExpression.FilterExpressionType) |  | Required; the type of this expression. |
| lhs | [FilterExpression](#capsule8.api.v0.FilterExpression) |  | Required for AND and OR types; this is the left hand side of the logical operator. |
| rhs | [FilterExpression](#capsule8.api.v0.FilterExpression) |  | Required for AND and OR types; this is the right hand side of the logical operator. |
| predicate | [FilterPredicate](#capsule8.api.v0.FilterPredicate) |  | Required for PREDICATE types; this is the comparison operation to perform. |






<a name="capsule8.api.v0.FilterPredicate"/>

### FilterPredicate
This represents a single comparison operation, which is of the form
&lt;field name&gt; &lt;predicate type&gt; &lt;value&gt;


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| type | [FilterPredicate.PredicateType](#capsule8.api.v0.FilterPredicate.PredicateType) |  | Requried; the comparison type of this predicate. |
| field_name | [string](#string) |  | Required; the name of the field to compare with. |
| value_type | [FilterPredicate.ValueType](#capsule8.api.v0.FilterPredicate.ValueType) |  | Required; the type of the value to compare against the value represented by the field. |
| signed_value | [sint64](#sint64) |  | Required when the value type is SIGNED; this is the actual signed value to use in the comparison. |
| unsigned_value | [uint64](#uint64) |  | Required when the value type is UNSIGNED; this is the actual unsigned value to use in the comparison. |
| string_value | [string](#string) |  | Required when the value type is STRING; this is the actual string value to use in the comparison. |






<a name="capsule8.api.v0.KernelFunctionCallFilter"/>

### KernelFunctionCallFilter
The KernelFunctionCallFilter specifies which kernel function call
events to include in the Subscription. The arguments map defines
values that will be fetched at each call and returned along with
the event. In order to minimize event volume, a filter may be
included that filters the kernel function calls based on the
observed values of the specified arguments at the time of the
kernel function call.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| type | [KernelFunctionCallEventType](#capsule8.api.v0.KernelFunctionCallEventType) |  | Required; the kernel function call event type to match |
| symbol | [string](#string) |  | Required; the kernel symbol to match on |
| arguments | [KernelFunctionCallFilter.ArgumentsEntry](#capsule8.api.v0.KernelFunctionCallFilter.ArgumentsEntry) | repeated | Optional; the field names and data to be returned by the kernel when the event triggers. Note that this is a map. The keys are the names to assign to the returned fields, and the values are a string describing the data to return, usually an expression involving the register containing the desired data and a suffix indicating the type of the data (e.g., &#34;s32&#34;, &#34;string&#34;, &#34;u64&#34;, etc.). This map is used to construct the &#34;fetchargs&#34; passed to the kernel when creating the kernel probe. |
| filter | [string](#string) |  | Optional; a filter to apply to kernel probe. This is applied to the kernel probe unmodified and should be in the format that the kernel understands. Note: this field will be removed post-0.1.0 in favor of a more flexible alternative. |






<a name="capsule8.api.v0.KernelFunctionCallFilter.ArgumentsEntry"/>

### KernelFunctionCallFilter.ArgumentsEntry



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| key | [string](#string) |  |  |
| value | [string](#string) |  |  |






<a name="capsule8.api.v0.LimitModifier"/>

### LimitModifier
The LimitModifier cancels the subscription on each Sensor after the
specified number of events. The entire Subscription may return more
events that this depending on how many active Sensors there are.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| limit | [int64](#int64) |  | Limit the number of events |






<a name="capsule8.api.v0.Modifier"/>

### Modifier
Modifier specifies which stream modifiers to apply if any. For a given
stream, a modifier can apply a throttle or limit etc. Modifiers can be
used together.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| throttle | [ThrottleModifier](#capsule8.api.v0.ThrottleModifier) |  |  |
| limit | [LimitModifier](#capsule8.api.v0.LimitModifier) |  |  |






<a name="capsule8.api.v0.NetworkEventFilter"/>

### NetworkEventFilter
The NetworkEventFilter specifies which network events to include in
the Subscription. The included filter can be used to specify
precisely which network events should be included.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| type | [NetworkEventType](#capsule8.api.v0.NetworkEventType) |  | Required; the network event type to match |
| filter | [FilterExpression](#capsule8.api.v0.FilterExpression) |  | Optional; a filter to apply to events. Only events for which the evaluation of the filter expression is true will be returned. Note: this field will be removed post-0.1.0 in favor of a more flexible alternative. |






<a name="capsule8.api.v0.ProcessEventFilter"/>

### ProcessEventFilter
The ProcessEventFilter specifies which process events to include in
the Subscription. The specified fields are effectively &#34;ANDed&#34; to
specify a matching event.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| type | [ProcessEventType](#capsule8.api.v0.ProcessEventType) |  | Required; the process event type to match |
| exec_filename | [.google.protobuf.StringValue](#capsule8.api.v0..google.protobuf.StringValue) |  | Optional; require exact match on the filename passed to execve(2) |
| exec_filename_pattern | [.google.protobuf.StringValue](#capsule8.api.v0..google.protobuf.StringValue) |  | Optional; require pattern match on the filename passed to execve(2) |
| exit_code | [.google.protobuf.Int32Value](#capsule8.api.v0..google.protobuf.Int32Value) |  | Optional; require exact match on exit code |






<a name="capsule8.api.v0.Subscription"/>

### Subscription
The Subscription message identifies a subscriber&#39;s interest in
telemetry events.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| event_filter | [EventFilter](#capsule8.api.v0.EventFilter) |  | Return events matching one or more of the specified event filters. If no event filters are specified, then no events will be returned. |
| container_filter | [ContainerFilter](#capsule8.api.v0.ContainerFilter) |  | If not empty, then only return events from containers matched by one or more of the specified container filters. |
| since_duration | [.google.protobuf.Int64Value](#capsule8.api.v0..google.protobuf.Int64Value) |  | If not empty, then only return events that occurred after the specified relative duration subtracted from the current time (recorder time). If the resulting time is in the past, then the subscription will search for historic events before streaming live ones. |
| for_duration | [.google.protobuf.Int64Value](#capsule8.api.v0..google.protobuf.Int64Value) |  | If not empty, then only return events that occurred before the specified relative duration added to `since_duration`. If `since_duration` is not supplied, return events from now and until the specified relative duration is hit. |
| modifier | [Modifier](#capsule8.api.v0.Modifier) |  | If not empty, apply the specified modifier to the subscription. |






<a name="capsule8.api.v0.SyscallEventFilter"/>

### SyscallEventFilter
The SyscallEventFilter specifies which system call events to
include in the Subscription. The specified fields are effectively
&#34;ANDed&#34; to specify a matching event.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| type | [SyscallEventType](#capsule8.api.v0.SyscallEventType) |  | Required; type of system call event (entry or exit) |
| id | [.google.protobuf.Int64Value](#capsule8.api.v0..google.protobuf.Int64Value) |  | Required; system call number from arch/x86/entry/syscalls/syscall_64.tbl |
| ret | [.google.protobuf.Int64Value](#capsule8.api.v0..google.protobuf.Int64Value) |  | Optional; return value of the system call (if type indicates exit). |






<a name="capsule8.api.v0.ThrottleModifier"/>

### ThrottleModifier
The ThrottleModifier modulates events sent by the Sensor to one per
time interval specified.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| interval | [int64](#int64) |  | Required; the interval to use |
| interval_type | [ThrottleModifier.IntervalType](#capsule8.api.v0.ThrottleModifier.IntervalType) |  | Required; the intreval type (milliseconds, seconds, etc.) |






<a name="capsule8.api.v0.TickerEventFilter"/>

### TickerEventFilter
The TickerEventFilter configures a ticker stream generator and
includes events from it in the Subscription.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| interval | [int64](#int64) |  | Required; the interval at which ticker events are generated |





 


<a name="capsule8.api.v0.ContainerEventView"/>

### ContainerEventView
The ContainerEventView specifies the level of detail to include for
ContainerEvents.

| Name | Number | Description |
| ---- | ------ | ----------- |
| BASIC | 0 | Default view of a ContainerEvent includes just basic information |
| FULL | 1 | Full view of a ContainerEvent includes raw Docker and OCI config JSON payloads |



<a name="capsule8.api.v0.FilterExpression.FilterExpressionType"/>

### FilterExpression.FilterExpressionType


| Name | Number | Description |
| ---- | ------ | ----------- |
| PREDICATE | 0 | PREDICATE is a comparison. When this type is specified for an expression, the predicate field must not be nil. |
| AND | 1 | AND is a logical combination of two expressions for which the lhs and rhs fields must not be nil, and both must evaluate true for this expression to evaluate true. |
| OR | 2 | OR is a logical combination of two expressions for which the lhs and rhs fields must not be nil, and only one must evaluate true for this expression to evaluate true. |



<a name="capsule8.api.v0.FilterPredicate.PredicateType"/>

### FilterPredicate.PredicateType
Possible types for FilterPredicate comparisons

| Name | Number | Description |
| ---- | ------ | ----------- |
| CONST | 0 | A constant value |
| EQ | 1 | Compare field equal to value |
| NE | 2 | Compare field not equal to value |
| LT | 3 | Compare field less than value |
| LE | 4 | Compare field less than or equal to value |
| GT | 5 | Compare field greater than value |
| GE | 6 | Compare field greater than or equal to value |
| GLOB | 7 | Compare field using wildcards (strings only; only leading and trailing &#39;*&#39; supported) |



<a name="capsule8.api.v0.FilterPredicate.ValueType"/>

### FilterPredicate.ValueType


| Name | Number | Description |
| ---- | ------ | ----------- |
| SIGNED | 0 | A signed integer value |
| UNSIGNED | 1 | An unsigned integer value |
| STRING | 2 | A string |



<a name="capsule8.api.v0.ThrottleModifier.IntervalType"/>

### ThrottleModifier.IntervalType
Possible interval types

| Name | Number | Description |
| ---- | ------ | ----------- |
| MILLISECOND | 0 | milliseconds |
| SECOND | 1 | seconds |
| MINUTE | 2 | minutes |
| HOUR | 3 | hours |


 

 

 



<a name="telemetry_service.proto"/>
<p align="right"><a href="#top">Top</a></p>

## telemetry_service.proto



<a name="capsule8.api.v0.GetEventsRequest"/>

### GetEventsRequest
A request message to initiate the streaming of telemetry events


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| subscription | [Subscription](#capsule8.api.v0.Subscription) |  | The Subscription message defines which events should be returned in the stream. |






<a name="capsule8.api.v0.GetEventsResponse"/>

### GetEventsResponse
A response message containing telemetry events


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| events | [TelemetryEvent](#capsule8.api.v0.TelemetryEvent) | repeated | Can publish one or more message(s) at a time |






<a name="capsule8.api.v0.TelemetryEvent"/>

### TelemetryEvent
A telemetry event received from a Sensor or Recorder.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| publish_time_micros | [int64](#int64) |  | The time that the event was received by the backplane (in micros since Unix epoch) |
| event | [Event](#capsule8.api.v0.Event) |  | The actual event observed by the Sensor. For historical event subscriptions, this event may be sent from the Recorder. |
| ack | [bytes](#bytes) |  | An opaque ack for the event. If present, this ack must be sent to the PubsubService&#39;s Acknowledge method or else the TelemetryService will re-transmit the event. |





 

 

 


<a name="capsule8.api.v0.TelemetryService"/>

### TelemetryService
Capsule8 Telemetry API

The Telemetry API allows you to subscribe to streams of live and
historical Events from Capsule8 Sensors.

| Method Name | Request Type | Response Type | Description |
| ----------- | ------------ | ------------- | ------------|
| GetEvents | [GetEventsRequest](#capsule8.api.v0.GetEventsRequest) | [GetEventsResponse](#capsule8.api.v0.GetEventsRequest) | Opens a new stream of telemetry events |

 



## Scalar Value Types

| .proto Type | Notes | C++ Type | Java Type | Python Type |
| ----------- | ----- | -------- | --------- | ----------- |
| <a name="double" /> double |  | double | double | float |
| <a name="float" /> float |  | float | float | float |
| <a name="int32" /> int32 | Uses variable-length encoding. Inefficient for encoding negative numbers – if your field is likely to have negative values, use sint32 instead. | int32 | int | int |
| <a name="int64" /> int64 | Uses variable-length encoding. Inefficient for encoding negative numbers – if your field is likely to have negative values, use sint64 instead. | int64 | long | int/long |
| <a name="uint32" /> uint32 | Uses variable-length encoding. | uint32 | int | int/long |
| <a name="uint64" /> uint64 | Uses variable-length encoding. | uint64 | long | int/long |
| <a name="sint32" /> sint32 | Uses variable-length encoding. Signed int value. These more efficiently encode negative numbers than regular int32s. | int32 | int | int |
| <a name="sint64" /> sint64 | Uses variable-length encoding. Signed int value. These more efficiently encode negative numbers than regular int64s. | int64 | long | int/long |
| <a name="fixed32" /> fixed32 | Always four bytes. More efficient than uint32 if values are often greater than 2^28. | uint32 | int | int |
| <a name="fixed64" /> fixed64 | Always eight bytes. More efficient than uint64 if values are often greater than 2^56. | uint64 | long | int/long |
| <a name="sfixed32" /> sfixed32 | Always four bytes. | int32 | int | int |
| <a name="sfixed64" /> sfixed64 | Always eight bytes. | int64 | long | int/long |
| <a name="bool" /> bool |  | bool | boolean | boolean |
| <a name="string" /> string | A string must always contain UTF-8 encoded or 7-bit ASCII text. | string | String | str/unicode |
| <a name="bytes" /> bytes | May contain any arbitrary sequence of bytes. | string | ByteString | str |

