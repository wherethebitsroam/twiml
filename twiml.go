package twiml

import (
	"encoding/xml"
	"io"
)

// NewResponse creates a new response object
func NewResponse() *Response {
	return &Response{}
}

// Add ...
func (r *Response) add(v interface{}) {
	r.Verbs = append(r.Verbs, v)
}

// Say ...
func (r *Response) Say(value string, attr *SayAttr) {
	r.add(&Say{
		Value:   value,
		SayAttr: *attr,
	})
}

// DialSimple ...
func (r *Response) DialSimple(value string, attr *DialAttr) {
	r.add(&Dial{
		Value:    value,
		DialAttr: *attr,
	})
}

// Dial ...
func (r *Response) Dial(attr *DialAttr) *Dial {
	d := &Dial{
		DialAttr: *attr,
	}
	r.add(d)
	return d
}

// Hangup ...
func (r *Response) Hangup() {
	r.add(&Hangup{})
}

// Pause ...
func (r *Response) Pause(attr *PauseAttr) {
	r.add(&Pause{
		PauseAttr: *attr,
	})
}

// Play ...
func (r *Response) Play(value string, attr *PlayAttr) {
	r.add(&Play{
		Value:    value,
		PlayAttr: *attr,
	})
}

// Gather ...
func (r *Response) Gather(attr *GatherAttr) *Gather {
	g := &Gather{
		GatherAttr: *attr,
	}
	r.add(g)
	return g
}

// ToXML ...
func (r *Response) ToXML(w io.Writer) error {
	enc := xml.NewEncoder(w)
	enc.Indent("", "  ")
	if err := enc.Encode(r); err != nil {
		return err
	}
	return nil
}

// Response ...
type Response struct {
	XMLName xml.Name `xml:"Response"`
	Verbs   []interface{}
}

// SayAttr ...
type SayAttr struct {
	Voice    string `xml:"voice,attr,omitempty"`
	Loop     int    `xml:"loop,attr,omitempty"`
	Language string `xml:"language,attr,omitempty"`
}

// Say ...
type Say struct {
	XMLName xml.Name `xml:"Say"`
	Value   string   `xml:",chardata"`
	SayAttr
}

// DialAttr ...
type DialAttr struct {
	Action       string `xml:"action,attr,omitempty"`
	Method       string `xml:"method,attr,omitempty"`
	Timeout      int    `xml:"timeout,attr,omitempty"`
	HangupOnStar bool   `xml:"hangupOnStar,attr,omitempty"`
	TimeLimit    int    `xml:"timeLimit,attr,omitempty"`
	CallerID     string `xml:"callerId,attr,omitempty"`
	Record       string `xml:"record,attr,omitempty"`
	Trim         string `xml:"trim,attr,omitempty"`
}

// Dial ...
type Dial struct {
	XMLName xml.Name `xml:"Dial"`
	Value   string   `xml:",chardata"`
	DialAttr
	Nouns []interface{}
}

// Number ...
func (d *Dial) Number(value string, attr *NumberAttr) {
	d.Nouns = append(d.Nouns, &Number{
		Value:      value,
		NumberAttr: *attr,
	})
}

// Client ...
func (d *Dial) Client(value string, attr *ClientAttr) {
	d.Nouns = append(d.Nouns, &Client{
		Value:      value,
		ClientAttr: *attr,
	})
}

// Sip ...
func (d *Dial) Sip(value string, attr *SipAttr) {
	d.Nouns = append(d.Nouns, &Sip{
		Value:   value,
		SipAttr: *attr,
	})
}

// RedirectAttr ...
type RedirectAttr struct {
	Method string `xml:"method,attr,omitempty"`
}

// Redirect ...
type Redirect struct {
	XMLName xml.Name `xml:"Redirect"`
	Value   string   `xml:",chardata"`
	RedirectAttr
}

// EnqueueAttr ...
type EnqueueAttr struct {
	Action      string `xml:"action,attr,omitempty"`
	Method      string `xml:"method,attr,omitempty"`
	Timeout     int    `xml:"timeout,attr,omitempty"`
	FinishOnKey string `xml:"finishOnKey,attr,omitempty"`
	NumDigits   int    `xml:"numDigits,attr,omitempty"`
}

// Enqueue ...
type Enqueue struct {
	XMLName xml.Name `xml:"Enqueue"`
}

// PauseAttr ...
type PauseAttr struct {
	Length int `xml:"length,attr,omitempty"`
}

// Pause ...
type Pause struct {
	XMLName xml.Name `xml:"Pause"`
	PauseAttr
}

// PlayAttr ...
type PlayAttr struct {
	Loop   int    `xml:"loop,attr,omitempty"`
	Digits string `xml:"digits,attr,omitempty"`
}

// Play ...
type Play struct {
	XMLName xml.Name `xml:"Play"`
	Value   string   `xml:",chardata"`
	PlayAttr
}

// GatherAttr ...
type GatherAttr struct {
	Action      string `xml:"action,attr,omitempty"`
	Method      string `xml:"method,attr,omitempty"`
	Timeout     int    `xml:"timeout,attr,omitempty"`
	FinishOnKey string `xml:"finishOnKey,attr,omitempty"`
	NumDigits   int    `xml:"numDigits,attr,omitempty"`
}

// Gather ...
type Gather struct {
	XMLName xml.Name `xml:"Gather"`
	Verbs   []interface{}
	GatherAttr
}

// Say ...
func (g *Gather) Say(value string, attr *SayAttr) {
	g.Verbs = append(g.Verbs, &Say{
		Value:   value,
		SayAttr: *attr,
	})
}

// Play ...
func (g *Gather) Play(value string, attr *PlayAttr) {
	g.Verbs = append(g.Verbs, &Play{
		Value:    value,
		PlayAttr: *attr,
	})
}

// Pause ...
func (g *Gather) Pause(attr *PauseAttr) {
	g.Verbs = append(g.Verbs, &Pause{
		PauseAttr: *attr,
	})
}

// Hangup ...
type Hangup struct {
	XMLName xml.Name `xml:"Hangup"`
}

// NumberAttr ...
type NumberAttr struct {
	SendDigits           string `xml:"sendDigits,attr,omitempty"`
	URL                  string `xml:"url,attr,omitempty"`
	Method               string `xml:"method,attr,omitempty"`
	StatusCallbackEvent  string `xml:"statusCallbackEvent,attr,omitempty"`
	StatusCallback       string `xml:"statusCallback,attr,omitempty"`
	StatusCallbackMethod string `xml:"statusCallbackMethod,attr,omitempty"`
}

// Number ...
type Number struct {
	XMLName xml.Name `xml:"Number"`
	Value   string   `xml:",chardata"`
	NumberAttr
}

// ClientAttr ...
type ClientAttr struct {
	URL                  string `xml:"url,attr,omitempty"`
	Method               string `xml:"method,attr,omitempty"`
	StatusCallbackEvent  string `xml:"statusCallbackEvent,attr,omitempty"`
	StatusCallback       string `xml:"statusCallback,attr,omitempty"`
	StatusCallbackMethod string `xml:"statusCallbackMethod,attr,omitempty"`
}

// Client ...
type Client struct {
	XMLName xml.Name `xml:"Client"`
	Value   string   `xml:",chardata"`
	ClientAttr
}

// SipAttr ...
type SipAttr struct {
	Username             string `xml:"username,attr,omitempty"`
	Password             string `xml:"password,attr,omitempty"`
	URL                  string `xml:"url,attr,omitempty"`
	Method               string `xml:"method,attr,omitempty"`
	StatusCallbackEvent  string `xml:"statusCallbackEvent,attr,omitempty"`
	StatusCallback       string `xml:"statusCallback,attr,omitempty"`
	StatusCallbackMethod string `xml:"statusCallbackMethod,attr,omitempty"`
}

// Sip ...
type Sip struct {
	XMLName xml.Name `xml:"Sip"`
	Value   string   `xml:",chardata"`
	SipAttr
}

// Conference ...
type Conference struct {
}

// Queue ...
type Queue struct {
}
