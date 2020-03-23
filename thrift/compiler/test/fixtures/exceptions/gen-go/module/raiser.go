// Autogenerated by Thrift Compiler (facebook)
// DO NOT EDIT UNLESS YOU ARE SURE THAT YOU KNOW WHAT YOU ARE DOING
// @generated

package module

import (
	"bytes"
	"context"
	"sync"
	"fmt"
	thrift "github.com/facebook/fbthrift/thrift/lib/go/thrift"
)

// (needed to ensure safety because of naive import list construction.)
var _ = thrift.ZERO
var _ = fmt.Printf
var _ = sync.Mutex{}
var _ = bytes.Equal
var _ = context.Background

type Raiser interface {
  DoBland() (err error)
  DoRaise() (err error)
  Get200() (_r string, err error)
  Get500() (_r string, err error)
}

type RaiserClientInterface interface {
  thrift.ClientInterface
  DoBland() (err error)
  DoRaise() (err error)
  Get200() (_r string, err error)
  Get500() (_r string, err error)
}

type RaiserClient struct {
  RaiserClientInterface
  CC thrift.ClientConn
}

func(client *RaiserClient) Open() error {
  return client.CC.Open()
}

func(client *RaiserClient) Close() error {
  return client.CC.Close()
}

func(client *RaiserClient) IsOpen() bool {
  return client.CC.IsOpen()
}

func NewRaiserClientFactory(t thrift.Transport, f thrift.ProtocolFactory) *RaiserClient {
  return &RaiserClient{ CC: thrift.NewClientConn(t, f) }
}

func NewRaiserClient(t thrift.Transport, iprot thrift.Protocol, oprot thrift.Protocol) *RaiserClient {
  return &RaiserClient{ CC: thrift.NewClientConnWithProtocols(t, iprot, oprot) }
}

func (p *RaiserClient) DoBland() (err error) {
  var args RaiserDoBlandArgs
  err = p.CC.SendMsg("doBland", &args, thrift.CALL)
  if err != nil { return }
  return p.recvDoBland()
}


func (p *RaiserClient) recvDoBland() (err error) {
  var result RaiserDoBlandResult
  return p.CC.RecvMsg("doBland", &result)
}

func (p *RaiserClient) DoRaise() (err error) {
  var args RaiserDoRaiseArgs
  err = p.CC.SendMsg("doRaise", &args, thrift.CALL)
  if err != nil { return }
  return p.recvDoRaise()
}


func (p *RaiserClient) recvDoRaise() (err error) {
  var result RaiserDoRaiseResult
  err = p.CC.RecvMsg("doRaise", &result)
  if err != nil { return }
  if result.B != nil {
    err = result.B
    return 
  } else if result.F != nil {
    err = result.F
    return 
  } else if result.S != nil {
    err = result.S
    return 
  }
  return nil
}

func (p *RaiserClient) Get200() (_r string, err error) {
  var args RaiserGet200Args
  err = p.CC.SendMsg("get200", &args, thrift.CALL)
  if err != nil { return }
  return p.recvGet200()
}


func (p *RaiserClient) recvGet200() (value string, err error) {
  var result RaiserGet200Result
  err = p.CC.RecvMsg("get200", &result)
  if err != nil { return }

  return result.GetSuccess(), nil
}

func (p *RaiserClient) Get500() (_r string, err error) {
  var args RaiserGet500Args
  err = p.CC.SendMsg("get500", &args, thrift.CALL)
  if err != nil { return }
  return p.recvGet500()
}


func (p *RaiserClient) recvGet500() (value string, err error) {
  var result RaiserGet500Result
  err = p.CC.RecvMsg("get500", &result)
  if err != nil { return }
  if result.F != nil {
    err = result.F
    return 
  } else if result.B != nil {
    err = result.B
    return 
  } else if result.S != nil {
    err = result.S
    return 
  }
  return result.GetSuccess(), nil
}


type RaiserThreadsafeClient struct {
  RaiserClientInterface
  CC thrift.ClientConn
  Mu sync.Mutex
}

func(client *RaiserThreadsafeClient) Open() error {
  client.Mu.Lock()
  defer client.Mu.Unlock()
  return client.CC.Open()
}

func(client *RaiserThreadsafeClient) Close() error {
  client.Mu.Lock()
  defer client.Mu.Unlock()
  return client.CC.Close()
}

func(client *RaiserThreadsafeClient) IsOpen() bool {
  client.Mu.Lock()
  defer client.Mu.Unlock()
  return client.CC.IsOpen()
}

func NewRaiserThreadsafeClientFactory(t thrift.Transport, f thrift.ProtocolFactory) *RaiserThreadsafeClient {
  return &RaiserThreadsafeClient{ CC: thrift.NewClientConn(t, f) }
}

func NewRaiserThreadsafeClient(t thrift.Transport, iprot thrift.Protocol, oprot thrift.Protocol) *RaiserThreadsafeClient {
  return &RaiserThreadsafeClient{ CC: thrift.NewClientConnWithProtocols(t, iprot, oprot) }
}

func (p *RaiserThreadsafeClient) DoBland() (err error) {
  p.Mu.Lock()
  defer p.Mu.Unlock()
  var args RaiserDoBlandArgs
  err = p.CC.SendMsg("doBland", &args, thrift.CALL)
  if err != nil { return }
  return p.recvDoBland()
}


func (p *RaiserThreadsafeClient) recvDoBland() (err error) {
  var result RaiserDoBlandResult
  return p.CC.RecvMsg("doBland", &result)
}

func (p *RaiserThreadsafeClient) DoRaise() (err error) {
  p.Mu.Lock()
  defer p.Mu.Unlock()
  var args RaiserDoRaiseArgs
  err = p.CC.SendMsg("doRaise", &args, thrift.CALL)
  if err != nil { return }
  return p.recvDoRaise()
}


func (p *RaiserThreadsafeClient) recvDoRaise() (err error) {
  var result RaiserDoRaiseResult
  err = p.CC.RecvMsg("doRaise", &result)
  if err != nil { return }
  if result.B != nil {
    err = result.B
    return 
  } else if result.F != nil {
    err = result.F
    return 
  } else if result.S != nil {
    err = result.S
    return 
  }
  return nil
}

func (p *RaiserThreadsafeClient) Get200() (_r string, err error) {
  p.Mu.Lock()
  defer p.Mu.Unlock()
  var args RaiserGet200Args
  err = p.CC.SendMsg("get200", &args, thrift.CALL)
  if err != nil { return }
  return p.recvGet200()
}


func (p *RaiserThreadsafeClient) recvGet200() (value string, err error) {
  var result RaiserGet200Result
  err = p.CC.RecvMsg("get200", &result)
  if err != nil { return }

  return result.GetSuccess(), nil
}

func (p *RaiserThreadsafeClient) Get500() (_r string, err error) {
  p.Mu.Lock()
  defer p.Mu.Unlock()
  var args RaiserGet500Args
  err = p.CC.SendMsg("get500", &args, thrift.CALL)
  if err != nil { return }
  return p.recvGet500()
}


func (p *RaiserThreadsafeClient) recvGet500() (value string, err error) {
  var result RaiserGet500Result
  err = p.CC.RecvMsg("get500", &result)
  if err != nil { return }
  if result.F != nil {
    err = result.F
    return 
  } else if result.B != nil {
    err = result.B
    return 
  } else if result.S != nil {
    err = result.S
    return 
  }
  return result.GetSuccess(), nil
}


type RaiserProcessor struct {
  processorMap map[string]thrift.ProcessorFunction
  handler Raiser
}

func (p *RaiserProcessor) AddToProcessorMap(key string, processor thrift.ProcessorFunction) {
  p.processorMap[key] = processor
}

func (p *RaiserProcessor) GetProcessorFunction(key string) (processor thrift.ProcessorFunction, err error) {
  if processor, ok := p.processorMap[key]; ok {
    return processor, nil
  }
  return nil, nil // generic error message will be sent
}

func (p *RaiserProcessor) ProcessorMap() map[string]thrift.ProcessorFunction {
  return p.processorMap
}

func NewRaiserProcessor(handler Raiser) *RaiserProcessor {
  self0 := &RaiserProcessor{handler:handler, processorMap:make(map[string]thrift.ProcessorFunction)}
  self0.processorMap["doBland"] = &raiserProcessorDoBland{handler:handler}
  self0.processorMap["doRaise"] = &raiserProcessorDoRaise{handler:handler}
  self0.processorMap["get200"] = &raiserProcessorGet200{handler:handler}
  self0.processorMap["get500"] = &raiserProcessorGet500{handler:handler}
  return self0
}

type raiserProcessorDoBland struct {
  handler Raiser
}

func (p *raiserProcessorDoBland) Read(iprot thrift.Protocol) (thrift.Struct, thrift.Exception) {
  args := RaiserDoBlandArgs{}
  if err := args.Read(iprot); err != nil {
    return nil, err
  }
  iprot.ReadMessageEnd()
  return &args, nil
}

func (p *raiserProcessorDoBland) Write(seqId int32, result thrift.WritableStruct, oprot thrift.Protocol) (err thrift.Exception) {
  var err2 error
  messageType := thrift.REPLY
  switch result.(type) {
  case thrift.ApplicationException:
    messageType = thrift.EXCEPTION
  }
  if err2 = oprot.WriteMessageBegin("doBland", messageType, seqId); err2 != nil {
    err = err2
  }
  if err2 = result.Write(oprot); err == nil && err2 != nil {
    err = err2
  }
  if err2 = oprot.WriteMessageEnd(); err == nil && err2 != nil {
    err = err2
  }
  if err2 = oprot.Flush(); err == nil && err2 != nil {
    err = err2
  }
  return err
}

func (p *raiserProcessorDoBland) Run(argStruct thrift.Struct) (thrift.WritableStruct, thrift.ApplicationException) {
  var result RaiserDoBlandResult
  if err := p.handler.DoBland(); err != nil {
    switch err.(type) {
    default:
      x := thrift.NewApplicationException(thrift.INTERNAL_ERROR, "Internal error processing doBland: " + err.Error())
      return x, x
    }
  }
  return &result, nil
}

type raiserProcessorDoRaise struct {
  handler Raiser
}

func (p *raiserProcessorDoRaise) Read(iprot thrift.Protocol) (thrift.Struct, thrift.Exception) {
  args := RaiserDoRaiseArgs{}
  if err := args.Read(iprot); err != nil {
    return nil, err
  }
  iprot.ReadMessageEnd()
  return &args, nil
}

func (p *raiserProcessorDoRaise) Write(seqId int32, result thrift.WritableStruct, oprot thrift.Protocol) (err thrift.Exception) {
  var err2 error
  messageType := thrift.REPLY
  switch v := result.(type) {
  case *Banal:
    msg := RaiserDoRaiseResult{B: v}
    result = &msg
  case *Fiery:
    msg := RaiserDoRaiseResult{F: v}
    result = &msg
  case *Serious:
    msg := RaiserDoRaiseResult{S: v}
    result = &msg
  case thrift.ApplicationException:
    messageType = thrift.EXCEPTION
  }
  if err2 = oprot.WriteMessageBegin("doRaise", messageType, seqId); err2 != nil {
    err = err2
  }
  if err2 = result.Write(oprot); err == nil && err2 != nil {
    err = err2
  }
  if err2 = oprot.WriteMessageEnd(); err == nil && err2 != nil {
    err = err2
  }
  if err2 = oprot.Flush(); err == nil && err2 != nil {
    err = err2
  }
  return err
}

func (p *raiserProcessorDoRaise) Run(argStruct thrift.Struct) (thrift.WritableStruct, thrift.ApplicationException) {
  var result RaiserDoRaiseResult
  if err := p.handler.DoRaise(); err != nil {
    switch v := err.(type) {
    case *Banal:
      result.B = v
    case *Fiery:
      result.F = v
    case *Serious:
      result.S = v
    default:
      x := thrift.NewApplicationException(thrift.INTERNAL_ERROR, "Internal error processing doRaise: " + err.Error())
      return x, x
    }
  }
  return &result, nil
}

type raiserProcessorGet200 struct {
  handler Raiser
}

func (p *raiserProcessorGet200) Read(iprot thrift.Protocol) (thrift.Struct, thrift.Exception) {
  args := RaiserGet200Args{}
  if err := args.Read(iprot); err != nil {
    return nil, err
  }
  iprot.ReadMessageEnd()
  return &args, nil
}

func (p *raiserProcessorGet200) Write(seqId int32, result thrift.WritableStruct, oprot thrift.Protocol) (err thrift.Exception) {
  var err2 error
  messageType := thrift.REPLY
  switch result.(type) {
  case thrift.ApplicationException:
    messageType = thrift.EXCEPTION
  }
  if err2 = oprot.WriteMessageBegin("get200", messageType, seqId); err2 != nil {
    err = err2
  }
  if err2 = result.Write(oprot); err == nil && err2 != nil {
    err = err2
  }
  if err2 = oprot.WriteMessageEnd(); err == nil && err2 != nil {
    err = err2
  }
  if err2 = oprot.Flush(); err == nil && err2 != nil {
    err = err2
  }
  return err
}

func (p *raiserProcessorGet200) Run(argStruct thrift.Struct) (thrift.WritableStruct, thrift.ApplicationException) {
  var result RaiserGet200Result
  if retval, err := p.handler.Get200(); err != nil {
    switch err.(type) {
    default:
      x := thrift.NewApplicationException(thrift.INTERNAL_ERROR, "Internal error processing get200: " + err.Error())
      return x, x
    }
  } else {
    result.Success = &retval
  }
  return &result, nil
}

type raiserProcessorGet500 struct {
  handler Raiser
}

func (p *raiserProcessorGet500) Read(iprot thrift.Protocol) (thrift.Struct, thrift.Exception) {
  args := RaiserGet500Args{}
  if err := args.Read(iprot); err != nil {
    return nil, err
  }
  iprot.ReadMessageEnd()
  return &args, nil
}

func (p *raiserProcessorGet500) Write(seqId int32, result thrift.WritableStruct, oprot thrift.Protocol) (err thrift.Exception) {
  var err2 error
  messageType := thrift.REPLY
  switch v := result.(type) {
  case *Fiery:
    msg := RaiserGet500Result{F: v}
    result = &msg
  case *Banal:
    msg := RaiserGet500Result{B: v}
    result = &msg
  case *Serious:
    msg := RaiserGet500Result{S: v}
    result = &msg
  case thrift.ApplicationException:
    messageType = thrift.EXCEPTION
  }
  if err2 = oprot.WriteMessageBegin("get500", messageType, seqId); err2 != nil {
    err = err2
  }
  if err2 = result.Write(oprot); err == nil && err2 != nil {
    err = err2
  }
  if err2 = oprot.WriteMessageEnd(); err == nil && err2 != nil {
    err = err2
  }
  if err2 = oprot.Flush(); err == nil && err2 != nil {
    err = err2
  }
  return err
}

func (p *raiserProcessorGet500) Run(argStruct thrift.Struct) (thrift.WritableStruct, thrift.ApplicationException) {
  var result RaiserGet500Result
  if retval, err := p.handler.Get500(); err != nil {
    switch v := err.(type) {
    case *Fiery:
      result.F = v
    case *Banal:
      result.B = v
    case *Serious:
      result.S = v
    default:
      x := thrift.NewApplicationException(thrift.INTERNAL_ERROR, "Internal error processing get500: " + err.Error())
      return x, x
    }
  } else {
    result.Success = &retval
  }
  return &result, nil
}


// HELPER FUNCTIONS AND STRUCTURES

type RaiserDoBlandArgs struct {
  thrift.IRequest
}

func NewRaiserDoBlandArgs() *RaiserDoBlandArgs {
  return &RaiserDoBlandArgs{}
}

func (p *RaiserDoBlandArgs) Read(iprot thrift.Protocol) error {
  if _, err := iprot.ReadStructBegin(); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T read error: ", p), err)
  }


  for {
    _, fieldTypeId, fieldId, err := iprot.ReadFieldBegin()
    if err != nil {
      return thrift.PrependError(fmt.Sprintf("%T field %d read error: ", p, fieldId), err)
    }
    if fieldTypeId == thrift.STOP { break; }
    if err := iprot.Skip(fieldTypeId); err != nil {
      return err
    }
    if err := iprot.ReadFieldEnd(); err != nil {
      return err
    }
  }
  if err := iprot.ReadStructEnd(); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T read struct end error: ", p), err)
  }
  return nil
}

func (p *RaiserDoBlandArgs) Write(oprot thrift.Protocol) error {
  if err := oprot.WriteStructBegin("doBland_args"); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T write struct begin error: ", p), err) }
  if err := oprot.WriteFieldStop(); err != nil {
    return thrift.PrependError("write field stop error: ", err) }
  if err := oprot.WriteStructEnd(); err != nil {
    return thrift.PrependError("write struct stop error: ", err) }
  return nil
}

func (p *RaiserDoBlandArgs) String() string {
  if p == nil {
    return "<nil>"
  }

  return fmt.Sprintf("RaiserDoBlandArgs({})")
}

type RaiserDoBlandResult struct {
  thrift.IResponse
}

func NewRaiserDoBlandResult() *RaiserDoBlandResult {
  return &RaiserDoBlandResult{}
}

func (p *RaiserDoBlandResult) Read(iprot thrift.Protocol) error {
  if _, err := iprot.ReadStructBegin(); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T read error: ", p), err)
  }


  for {
    _, fieldTypeId, fieldId, err := iprot.ReadFieldBegin()
    if err != nil {
      return thrift.PrependError(fmt.Sprintf("%T field %d read error: ", p, fieldId), err)
    }
    if fieldTypeId == thrift.STOP { break; }
    if err := iprot.Skip(fieldTypeId); err != nil {
      return err
    }
    if err := iprot.ReadFieldEnd(); err != nil {
      return err
    }
  }
  if err := iprot.ReadStructEnd(); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T read struct end error: ", p), err)
  }
  return nil
}

func (p *RaiserDoBlandResult) Write(oprot thrift.Protocol) error {
  if err := oprot.WriteStructBegin("doBland_result"); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T write struct begin error: ", p), err) }
  if err := oprot.WriteFieldStop(); err != nil {
    return thrift.PrependError("write field stop error: ", err) }
  if err := oprot.WriteStructEnd(); err != nil {
    return thrift.PrependError("write struct stop error: ", err) }
  return nil
}

func (p *RaiserDoBlandResult) String() string {
  if p == nil {
    return "<nil>"
  }

  return fmt.Sprintf("RaiserDoBlandResult({})")
}

type RaiserDoRaiseArgs struct {
  thrift.IRequest
}

func NewRaiserDoRaiseArgs() *RaiserDoRaiseArgs {
  return &RaiserDoRaiseArgs{}
}

func (p *RaiserDoRaiseArgs) Read(iprot thrift.Protocol) error {
  if _, err := iprot.ReadStructBegin(); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T read error: ", p), err)
  }


  for {
    _, fieldTypeId, fieldId, err := iprot.ReadFieldBegin()
    if err != nil {
      return thrift.PrependError(fmt.Sprintf("%T field %d read error: ", p, fieldId), err)
    }
    if fieldTypeId == thrift.STOP { break; }
    if err := iprot.Skip(fieldTypeId); err != nil {
      return err
    }
    if err := iprot.ReadFieldEnd(); err != nil {
      return err
    }
  }
  if err := iprot.ReadStructEnd(); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T read struct end error: ", p), err)
  }
  return nil
}

func (p *RaiserDoRaiseArgs) Write(oprot thrift.Protocol) error {
  if err := oprot.WriteStructBegin("doRaise_args"); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T write struct begin error: ", p), err) }
  if err := oprot.WriteFieldStop(); err != nil {
    return thrift.PrependError("write field stop error: ", err) }
  if err := oprot.WriteStructEnd(); err != nil {
    return thrift.PrependError("write struct stop error: ", err) }
  return nil
}

func (p *RaiserDoRaiseArgs) String() string {
  if p == nil {
    return "<nil>"
  }

  return fmt.Sprintf("RaiserDoRaiseArgs({})")
}

// Attributes:
//  - B
//  - F
//  - S
type RaiserDoRaiseResult struct {
  thrift.IResponse
  B *Banal `thrift:"b,1" db:"b" json:"b,omitempty"`
  F *Fiery `thrift:"f,2" db:"f" json:"f,omitempty"`
  S *Serious `thrift:"s,3" db:"s" json:"s,omitempty"`
}

func NewRaiserDoRaiseResult() *RaiserDoRaiseResult {
  return &RaiserDoRaiseResult{}
}

var RaiserDoRaiseResult_B_DEFAULT *Banal
func (p *RaiserDoRaiseResult) GetB() *Banal {
  if !p.IsSetB() {
    return RaiserDoRaiseResult_B_DEFAULT
  }
return p.B
}
var RaiserDoRaiseResult_F_DEFAULT *Fiery
func (p *RaiserDoRaiseResult) GetF() *Fiery {
  if !p.IsSetF() {
    return RaiserDoRaiseResult_F_DEFAULT
  }
return p.F
}
var RaiserDoRaiseResult_S_DEFAULT *Serious
func (p *RaiserDoRaiseResult) GetS() *Serious {
  if !p.IsSetS() {
    return RaiserDoRaiseResult_S_DEFAULT
  }
return p.S
}
func (p *RaiserDoRaiseResult) IsSetB() bool {
  return p.B != nil
}

func (p *RaiserDoRaiseResult) IsSetF() bool {
  return p.F != nil
}

func (p *RaiserDoRaiseResult) IsSetS() bool {
  return p.S != nil
}

func (p *RaiserDoRaiseResult) Read(iprot thrift.Protocol) error {
  if _, err := iprot.ReadStructBegin(); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T read error: ", p), err)
  }


  for {
    _, fieldTypeId, fieldId, err := iprot.ReadFieldBegin()
    if err != nil {
      return thrift.PrependError(fmt.Sprintf("%T field %d read error: ", p, fieldId), err)
    }
    if fieldTypeId == thrift.STOP { break; }
    switch fieldId {
    case 1:
      if err := p.ReadField1(iprot); err != nil {
        return err
      }
    case 2:
      if err := p.ReadField2(iprot); err != nil {
        return err
      }
    case 3:
      if err := p.ReadField3(iprot); err != nil {
        return err
      }
    default:
      if err := iprot.Skip(fieldTypeId); err != nil {
        return err
      }
    }
    if err := iprot.ReadFieldEnd(); err != nil {
      return err
    }
  }
  if err := iprot.ReadStructEnd(); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T read struct end error: ", p), err)
  }
  return nil
}

func (p *RaiserDoRaiseResult)  ReadField1(iprot thrift.Protocol) error {
  p.B = NewBanal()
  if err := p.B.Read(iprot); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T error reading struct: ", p.B), err)
  }
  return nil
}

func (p *RaiserDoRaiseResult)  ReadField2(iprot thrift.Protocol) error {
  p.F = NewFiery()
  if err := p.F.Read(iprot); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T error reading struct: ", p.F), err)
  }
  return nil
}

func (p *RaiserDoRaiseResult)  ReadField3(iprot thrift.Protocol) error {
  p.S = NewSerious()
  if err := p.S.Read(iprot); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T error reading struct: ", p.S), err)
  }
  return nil
}

func (p *RaiserDoRaiseResult) Write(oprot thrift.Protocol) error {
  if err := oprot.WriteStructBegin("doRaise_result"); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T write struct begin error: ", p), err) }
  if err := p.writeField1(oprot); err != nil { return err }
  if err := p.writeField2(oprot); err != nil { return err }
  if err := p.writeField3(oprot); err != nil { return err }
  if err := oprot.WriteFieldStop(); err != nil {
    return thrift.PrependError("write field stop error: ", err) }
  if err := oprot.WriteStructEnd(); err != nil {
    return thrift.PrependError("write struct stop error: ", err) }
  return nil
}

func (p *RaiserDoRaiseResult) writeField1(oprot thrift.Protocol) (err error) {
  if p.IsSetB() {
    if err := oprot.WriteFieldBegin("b", thrift.STRUCT, 1); err != nil {
      return thrift.PrependError(fmt.Sprintf("%T write field begin error 1:b: ", p), err) }
    if err := p.B.Write(oprot); err != nil {
      return thrift.PrependError(fmt.Sprintf("%T error writing struct: ", p.B), err)
    }
    if err := oprot.WriteFieldEnd(); err != nil {
      return thrift.PrependError(fmt.Sprintf("%T write field end error 1:b: ", p), err) }
  }
  return err
}

func (p *RaiserDoRaiseResult) writeField2(oprot thrift.Protocol) (err error) {
  if p.IsSetF() {
    if err := oprot.WriteFieldBegin("f", thrift.STRUCT, 2); err != nil {
      return thrift.PrependError(fmt.Sprintf("%T write field begin error 2:f: ", p), err) }
    if err := p.F.Write(oprot); err != nil {
      return thrift.PrependError(fmt.Sprintf("%T error writing struct: ", p.F), err)
    }
    if err := oprot.WriteFieldEnd(); err != nil {
      return thrift.PrependError(fmt.Sprintf("%T write field end error 2:f: ", p), err) }
  }
  return err
}

func (p *RaiserDoRaiseResult) writeField3(oprot thrift.Protocol) (err error) {
  if p.IsSetS() {
    if err := oprot.WriteFieldBegin("s", thrift.STRUCT, 3); err != nil {
      return thrift.PrependError(fmt.Sprintf("%T write field begin error 3:s: ", p), err) }
    if err := p.S.Write(oprot); err != nil {
      return thrift.PrependError(fmt.Sprintf("%T error writing struct: ", p.S), err)
    }
    if err := oprot.WriteFieldEnd(); err != nil {
      return thrift.PrependError(fmt.Sprintf("%T write field end error 3:s: ", p), err) }
  }
  return err
}

func (p *RaiserDoRaiseResult) String() string {
  if p == nil {
    return "<nil>"
  }

  var bVal string
  if p.B == nil {
    bVal = "<nil>"
  }else{
    bVal = fmt.Sprint(*p.B)
  }
  var fVal string
  if p.F == nil {
    fVal = "<nil>"
  }else{
    fVal = fmt.Sprint(*p.F)
  }
  var sVal string
  if p.S == nil {
    sVal = "<nil>"
  }else{
    sVal = fmt.Sprint(*p.S)
  }
  return fmt.Sprintf("RaiserDoRaiseResult({B:%s F:%s S:%s })", bVal, fVal, sVal)
}

type RaiserGet200Args struct {
  thrift.IRequest
}

func NewRaiserGet200Args() *RaiserGet200Args {
  return &RaiserGet200Args{}
}

func (p *RaiserGet200Args) Read(iprot thrift.Protocol) error {
  if _, err := iprot.ReadStructBegin(); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T read error: ", p), err)
  }


  for {
    _, fieldTypeId, fieldId, err := iprot.ReadFieldBegin()
    if err != nil {
      return thrift.PrependError(fmt.Sprintf("%T field %d read error: ", p, fieldId), err)
    }
    if fieldTypeId == thrift.STOP { break; }
    if err := iprot.Skip(fieldTypeId); err != nil {
      return err
    }
    if err := iprot.ReadFieldEnd(); err != nil {
      return err
    }
  }
  if err := iprot.ReadStructEnd(); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T read struct end error: ", p), err)
  }
  return nil
}

func (p *RaiserGet200Args) Write(oprot thrift.Protocol) error {
  if err := oprot.WriteStructBegin("get200_args"); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T write struct begin error: ", p), err) }
  if err := oprot.WriteFieldStop(); err != nil {
    return thrift.PrependError("write field stop error: ", err) }
  if err := oprot.WriteStructEnd(); err != nil {
    return thrift.PrependError("write struct stop error: ", err) }
  return nil
}

func (p *RaiserGet200Args) String() string {
  if p == nil {
    return "<nil>"
  }

  return fmt.Sprintf("RaiserGet200Args({})")
}

// Attributes:
//  - Success
type RaiserGet200Result struct {
  thrift.IResponse
  Success *string `thrift:"success,0" db:"success" json:"success,omitempty"`
}

func NewRaiserGet200Result() *RaiserGet200Result {
  return &RaiserGet200Result{}
}

var RaiserGet200Result_Success_DEFAULT string
func (p *RaiserGet200Result) GetSuccess() string {
  if !p.IsSetSuccess() {
    return RaiserGet200Result_Success_DEFAULT
  }
return *p.Success
}
func (p *RaiserGet200Result) IsSetSuccess() bool {
  return p.Success != nil
}

func (p *RaiserGet200Result) Read(iprot thrift.Protocol) error {
  if _, err := iprot.ReadStructBegin(); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T read error: ", p), err)
  }


  for {
    _, fieldTypeId, fieldId, err := iprot.ReadFieldBegin()
    if err != nil {
      return thrift.PrependError(fmt.Sprintf("%T field %d read error: ", p, fieldId), err)
    }
    if fieldTypeId == thrift.STOP { break; }
    switch fieldId {
    case 0:
      if err := p.ReadField0(iprot); err != nil {
        return err
      }
    default:
      if err := iprot.Skip(fieldTypeId); err != nil {
        return err
      }
    }
    if err := iprot.ReadFieldEnd(); err != nil {
      return err
    }
  }
  if err := iprot.ReadStructEnd(); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T read struct end error: ", p), err)
  }
  return nil
}

func (p *RaiserGet200Result)  ReadField0(iprot thrift.Protocol) error {
  if v, err := iprot.ReadString(); err != nil {
  return thrift.PrependError("error reading field 0: ", err)
} else {
  p.Success = &v
}
  return nil
}

func (p *RaiserGet200Result) Write(oprot thrift.Protocol) error {
  if err := oprot.WriteStructBegin("get200_result"); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T write struct begin error: ", p), err) }
  if err := p.writeField0(oprot); err != nil { return err }
  if err := oprot.WriteFieldStop(); err != nil {
    return thrift.PrependError("write field stop error: ", err) }
  if err := oprot.WriteStructEnd(); err != nil {
    return thrift.PrependError("write struct stop error: ", err) }
  return nil
}

func (p *RaiserGet200Result) writeField0(oprot thrift.Protocol) (err error) {
  if p.IsSetSuccess() {
    if err := oprot.WriteFieldBegin("success", thrift.STRING, 0); err != nil {
      return thrift.PrependError(fmt.Sprintf("%T write field begin error 0:success: ", p), err) }
    if err := oprot.WriteString(string(*p.Success)); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T.success (0) field write error: ", p), err) }
    if err := oprot.WriteFieldEnd(); err != nil {
      return thrift.PrependError(fmt.Sprintf("%T write field end error 0:success: ", p), err) }
  }
  return err
}

func (p *RaiserGet200Result) String() string {
  if p == nil {
    return "<nil>"
  }

  var successVal string
  if p.Success == nil {
    successVal = "<nil>"
  }else{
    successVal = fmt.Sprint(*p.Success)
  }
  return fmt.Sprintf("RaiserGet200Result({Success:%s })", successVal)
}

type RaiserGet500Args struct {
  thrift.IRequest
}

func NewRaiserGet500Args() *RaiserGet500Args {
  return &RaiserGet500Args{}
}

func (p *RaiserGet500Args) Read(iprot thrift.Protocol) error {
  if _, err := iprot.ReadStructBegin(); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T read error: ", p), err)
  }


  for {
    _, fieldTypeId, fieldId, err := iprot.ReadFieldBegin()
    if err != nil {
      return thrift.PrependError(fmt.Sprintf("%T field %d read error: ", p, fieldId), err)
    }
    if fieldTypeId == thrift.STOP { break; }
    if err := iprot.Skip(fieldTypeId); err != nil {
      return err
    }
    if err := iprot.ReadFieldEnd(); err != nil {
      return err
    }
  }
  if err := iprot.ReadStructEnd(); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T read struct end error: ", p), err)
  }
  return nil
}

func (p *RaiserGet500Args) Write(oprot thrift.Protocol) error {
  if err := oprot.WriteStructBegin("get500_args"); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T write struct begin error: ", p), err) }
  if err := oprot.WriteFieldStop(); err != nil {
    return thrift.PrependError("write field stop error: ", err) }
  if err := oprot.WriteStructEnd(); err != nil {
    return thrift.PrependError("write struct stop error: ", err) }
  return nil
}

func (p *RaiserGet500Args) String() string {
  if p == nil {
    return "<nil>"
  }

  return fmt.Sprintf("RaiserGet500Args({})")
}

// Attributes:
//  - Success
//  - F
//  - B
//  - S
type RaiserGet500Result struct {
  thrift.IResponse
  Success *string `thrift:"success,0" db:"success" json:"success,omitempty"`
  F *Fiery `thrift:"f,1" db:"f" json:"f,omitempty"`
  B *Banal `thrift:"b,2" db:"b" json:"b,omitempty"`
  S *Serious `thrift:"s,3" db:"s" json:"s,omitempty"`
}

func NewRaiserGet500Result() *RaiserGet500Result {
  return &RaiserGet500Result{}
}

var RaiserGet500Result_Success_DEFAULT string
func (p *RaiserGet500Result) GetSuccess() string {
  if !p.IsSetSuccess() {
    return RaiserGet500Result_Success_DEFAULT
  }
return *p.Success
}
var RaiserGet500Result_F_DEFAULT *Fiery
func (p *RaiserGet500Result) GetF() *Fiery {
  if !p.IsSetF() {
    return RaiserGet500Result_F_DEFAULT
  }
return p.F
}
var RaiserGet500Result_B_DEFAULT *Banal
func (p *RaiserGet500Result) GetB() *Banal {
  if !p.IsSetB() {
    return RaiserGet500Result_B_DEFAULT
  }
return p.B
}
var RaiserGet500Result_S_DEFAULT *Serious
func (p *RaiserGet500Result) GetS() *Serious {
  if !p.IsSetS() {
    return RaiserGet500Result_S_DEFAULT
  }
return p.S
}
func (p *RaiserGet500Result) IsSetSuccess() bool {
  return p.Success != nil
}

func (p *RaiserGet500Result) IsSetF() bool {
  return p.F != nil
}

func (p *RaiserGet500Result) IsSetB() bool {
  return p.B != nil
}

func (p *RaiserGet500Result) IsSetS() bool {
  return p.S != nil
}

func (p *RaiserGet500Result) Read(iprot thrift.Protocol) error {
  if _, err := iprot.ReadStructBegin(); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T read error: ", p), err)
  }


  for {
    _, fieldTypeId, fieldId, err := iprot.ReadFieldBegin()
    if err != nil {
      return thrift.PrependError(fmt.Sprintf("%T field %d read error: ", p, fieldId), err)
    }
    if fieldTypeId == thrift.STOP { break; }
    switch fieldId {
    case 0:
      if err := p.ReadField0(iprot); err != nil {
        return err
      }
    case 1:
      if err := p.ReadField1(iprot); err != nil {
        return err
      }
    case 2:
      if err := p.ReadField2(iprot); err != nil {
        return err
      }
    case 3:
      if err := p.ReadField3(iprot); err != nil {
        return err
      }
    default:
      if err := iprot.Skip(fieldTypeId); err != nil {
        return err
      }
    }
    if err := iprot.ReadFieldEnd(); err != nil {
      return err
    }
  }
  if err := iprot.ReadStructEnd(); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T read struct end error: ", p), err)
  }
  return nil
}

func (p *RaiserGet500Result)  ReadField0(iprot thrift.Protocol) error {
  if v, err := iprot.ReadString(); err != nil {
  return thrift.PrependError("error reading field 0: ", err)
} else {
  p.Success = &v
}
  return nil
}

func (p *RaiserGet500Result)  ReadField1(iprot thrift.Protocol) error {
  p.F = NewFiery()
  if err := p.F.Read(iprot); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T error reading struct: ", p.F), err)
  }
  return nil
}

func (p *RaiserGet500Result)  ReadField2(iprot thrift.Protocol) error {
  p.B = NewBanal()
  if err := p.B.Read(iprot); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T error reading struct: ", p.B), err)
  }
  return nil
}

func (p *RaiserGet500Result)  ReadField3(iprot thrift.Protocol) error {
  p.S = NewSerious()
  if err := p.S.Read(iprot); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T error reading struct: ", p.S), err)
  }
  return nil
}

func (p *RaiserGet500Result) Write(oprot thrift.Protocol) error {
  if err := oprot.WriteStructBegin("get500_result"); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T write struct begin error: ", p), err) }
  if err := p.writeField0(oprot); err != nil { return err }
  if err := p.writeField1(oprot); err != nil { return err }
  if err := p.writeField2(oprot); err != nil { return err }
  if err := p.writeField3(oprot); err != nil { return err }
  if err := oprot.WriteFieldStop(); err != nil {
    return thrift.PrependError("write field stop error: ", err) }
  if err := oprot.WriteStructEnd(); err != nil {
    return thrift.PrependError("write struct stop error: ", err) }
  return nil
}

func (p *RaiserGet500Result) writeField0(oprot thrift.Protocol) (err error) {
  if p.IsSetSuccess() {
    if err := oprot.WriteFieldBegin("success", thrift.STRING, 0); err != nil {
      return thrift.PrependError(fmt.Sprintf("%T write field begin error 0:success: ", p), err) }
    if err := oprot.WriteString(string(*p.Success)); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T.success (0) field write error: ", p), err) }
    if err := oprot.WriteFieldEnd(); err != nil {
      return thrift.PrependError(fmt.Sprintf("%T write field end error 0:success: ", p), err) }
  }
  return err
}

func (p *RaiserGet500Result) writeField1(oprot thrift.Protocol) (err error) {
  if p.IsSetF() {
    if err := oprot.WriteFieldBegin("f", thrift.STRUCT, 1); err != nil {
      return thrift.PrependError(fmt.Sprintf("%T write field begin error 1:f: ", p), err) }
    if err := p.F.Write(oprot); err != nil {
      return thrift.PrependError(fmt.Sprintf("%T error writing struct: ", p.F), err)
    }
    if err := oprot.WriteFieldEnd(); err != nil {
      return thrift.PrependError(fmt.Sprintf("%T write field end error 1:f: ", p), err) }
  }
  return err
}

func (p *RaiserGet500Result) writeField2(oprot thrift.Protocol) (err error) {
  if p.IsSetB() {
    if err := oprot.WriteFieldBegin("b", thrift.STRUCT, 2); err != nil {
      return thrift.PrependError(fmt.Sprintf("%T write field begin error 2:b: ", p), err) }
    if err := p.B.Write(oprot); err != nil {
      return thrift.PrependError(fmt.Sprintf("%T error writing struct: ", p.B), err)
    }
    if err := oprot.WriteFieldEnd(); err != nil {
      return thrift.PrependError(fmt.Sprintf("%T write field end error 2:b: ", p), err) }
  }
  return err
}

func (p *RaiserGet500Result) writeField3(oprot thrift.Protocol) (err error) {
  if p.IsSetS() {
    if err := oprot.WriteFieldBegin("s", thrift.STRUCT, 3); err != nil {
      return thrift.PrependError(fmt.Sprintf("%T write field begin error 3:s: ", p), err) }
    if err := p.S.Write(oprot); err != nil {
      return thrift.PrependError(fmt.Sprintf("%T error writing struct: ", p.S), err)
    }
    if err := oprot.WriteFieldEnd(); err != nil {
      return thrift.PrependError(fmt.Sprintf("%T write field end error 3:s: ", p), err) }
  }
  return err
}

func (p *RaiserGet500Result) String() string {
  if p == nil {
    return "<nil>"
  }

  var successVal string
  if p.Success == nil {
    successVal = "<nil>"
  }else{
    successVal = fmt.Sprint(*p.Success)
  }
  var fVal string
  if p.F == nil {
    fVal = "<nil>"
  }else{
    fVal = fmt.Sprint(*p.F)
  }
  var bVal string
  if p.B == nil {
    bVal = "<nil>"
  }else{
    bVal = fmt.Sprint(*p.B)
  }
  var sVal string
  if p.S == nil {
    sVal = "<nil>"
  }else{
    sVal = fmt.Sprint(*p.S)
  }
  return fmt.Sprintf("RaiserGet500Result({Success:%s F:%s B:%s S:%s })", successVal, fVal, bVal, sVal)
}


