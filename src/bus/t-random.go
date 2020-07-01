package bus

import (
	"errors"
	"fmt"
	"math/rand"
	"strconv"
	"time"
)


func init() {
  InstallBus("random", &bus_random_ct{})
}


func _parse_random_slot(s string) (*bus_random_sl, error) {
  var t int
  var port int
  var tp SlotType

  n, err := fmt.Sscanf(s, "%c#%d", &t, &port)
  if err != nil {
    return nil, err
  }
  if n != 2 {
    return nil, errors.New("无效的slot格式")
  }
  switch (t) {
  case 'D':
    tp = SlotData
  case 'C':
    tp = SlotCtrl
  default:
    return nil, errors.New("无效的类型字符")
  }
  return &bus_random_sl{port, tp}, nil
}


type bus_random_ct struct {
}


func (*bus_random_ct) Name() string {
  return "随机数(测试)"
}


func (*bus_random_ct) Create(i *BusInfo) (Bus, error) {
  return &random_bus{i, BusStateStartup}, nil
}


// 接受任何字符串作为 slot
func (*bus_random_ct) ParseSlot(s string) (Slot, error) {
  return _parse_random_slot(s)
}


func (*bus_random_ct) SlotDesc(s string) (string, error) {
  slot, err := _parse_random_slot(s)
  if err != nil {
    return "", err
  }
  return slot.Desc(), nil
}


type bus_random_sl struct {
  port int
  tp SlotType
}


func (s *bus_random_sl) String() string {
  if s.tp == SlotData {
    return "D#"+ strconv.Itoa(s.port)
  } else {
    return "C#"+ strconv.Itoa(s.port)
  }
}


func (s *bus_random_sl) Desc() string {
  if s.tp == SlotData {
    return "虚拟数据 "+ strconv.Itoa(s.port)
  } else {
    return "虚拟控制 "+ strconv.Itoa(s.port)
  }
}


func (s *bus_random_sl) Type() SlotType {
  return s.tp
}


type random_bus struct {
  info    *BusInfo
  state   BusState
}


func (r *random_bus) start() error {
  r.state = BusStateSleep
  r.info.Tk.Start(func() {
    r.state = BusStateTask
    t := time.Now()
    
    for _, s := range r.info.SlotConf {
      d := IntData{rand.Int()}
      r.info.Recv.OnData(s, &t, &d)
    }

    r.state = BusStateSleep
  }, func() {
    r.stop()
  })
  return nil
}


func (r *random_bus) stop() {
  r.state = BusStateStop
}


func (*random_bus) SendCtrl(s Slot, d DataWrap) error {
  return nil
}


func (r *random_bus) State() BusState {
  return r.state
}
