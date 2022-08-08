package benchmark

import "time"

type SmallPayload struct {
	St   int    `json:"st"`
	Sid  int    `json:"sid"`
	Tt   string `json:"tt"`
	Gr   int    `json:"gr"`
	Uuid string `json:"uuid"`
	Ip   string `json:"ip"`
	Ua   string `json:"ua"`
	Tz   int    `json:"tz"`
	V    int    `json:"v"`
}

/*
   Medium payload (based on Clearbit API response)
*/
type CBAvatar struct {
	Url string `json:"url"`
}

type Avatars []*CBAvatar

type CBGravatar struct {
	Avatars Avatars `json:"avatars"`
}

type CBGithub struct {
	Followers int `json:"followers"`
}

type CBName struct {
	FullName string `json:"fullName"`
}

type CBPerson struct {
	Name     *CBName     `json:"name"`
	Github   *CBGithub   `json:"github"`
	Gravatar *CBGravatar `json:"gravatar"`
}

type MediumPayload struct {
	Person  *CBPerson `json:"person,omitempty"`
	Company string    `json:"company,omitempty"`
}

type DSUser struct {
	Username string `json:"username"`
}

type DSTopic struct {
	Id   int    `json:"id"`
	Slug string `json:"slug"`
}

type DSTopics []*DSTopic

type DSTopicsList struct {
	Topics        DSTopics `json:"topics"`
	MoreTopicsUrl string   `json:"more_topics_url"`
}

type DSUsers []*DSUser

type LargePayload struct {
	Users  DSUsers       `json:"users"`
	Topics *DSTopicsList `json:"topics,omitempty"`
}

type TestStruct struct {
	A string `json:"a,escape"`
}

type TestLargeStruct struct {
	Timestamp   time.Time   `json:"@timestamp"`
	Metadata    Metadata    `json:"@metadata"`
	Ecs         Ecs         `json:"ecs"`
	Host        Host        `json:"host"`
	Server      Server      `json:"server"`
	Status      string      `json:"status"`
	Source      Source      `json:"source"`
	Method      string      `json:"method"`
	HTTP        HTTP        `json:"http"`
	Network     Network     `json:"network"`
	URL         URL         `json:"url"`
	Client      Client      `json:"client"`
	Event       Event       `json:"event"`
	Query       string      `json:"query"`
	UserAgent   UserAgent   `json:"user_agent"`
	Destination Destination `json:"destination"`
	Type        string      `json:"type"`
	Agent       Agent       `json:"agent"`
}
type Metadata struct {
	Beat    string `json:"beat"`
	Type    string `json:"type"`
	Version string `json:"version"`
	Topic   string `json:"topic"`
}
type Ecs struct {
	Version string `json:"version"`
}
type Host struct {
	Name string `json:"name"`
}
type Server struct {
	IP     string `json:"ip"`
	Port   int    `json:"port"`
	Domain string `json:"domain"`
	Bytes  int    `json:"bytes"`
}
type Source struct {
	Bytes int    `json:"bytes"`
	IP    string `json:"ip"`
	Port  int    `json:"port"`
}
type Body struct {
	Content string `json:"content,escape"`
	Bytes   int    `json:"bytes"`
}
type ResponseHeaders struct {
	ContentLength    int    `json:"content-length"`
	TransferEncoding string `json:"transfer-encoding"`
	Connection       string `json:"connection"`
	CacheControl     string `json:"cache-control"`
	Pragma           string `json:"pragma"`
	Server           string `json:"server"`
	Date             string `json:"date"`
}
type Response struct {
	StatusCode   int             `json:"status_code"`
	Body         Body            `json:"body"`
	Bytes        int             `json:"bytes"`
	Headers      ResponseHeaders `json:"headers"`
	StatusPhrase string          `json:"status_phrase"`
}
type RequestHeaders struct {
	Referer        string `json:"referer"`
	XRequestedWith string `json:"x-requested-with"`
	YzClientIP     string `json:"yz_client_ip"`
	UserAgent      string `json:"user-agent"`
	AcceptLanguage string `json:"accept-language"`
	ContentLength  int    `json:"content-length"`
	XRealIP        string `json:"x-real-ip"`
	Pragma         string `json:"pragma"`
	Connection     string `json:"connection"`
	Accept         string `json:"accept"`
	Host           string `json:"host"`
	XForwardedFor  string `json:"x-forwarded-for"`
	ContentType    string `json:"content-type"`
}
type Request struct {
	Referrer string         `json:"referrer"`
	Bytes    int            `json:"bytes"`
	Headers  RequestHeaders `json:"headers"`
	Method   string         `json:"method"`
	Body     Body           `json:"body"`
}
type HTTP struct {
	Response Response `json:"response"`
	Version  string   `json:"version"`
	Request  Request  `json:"request"`
}
type Network struct {
	Type        string `json:"type"`
	Transport   string `json:"transport"`
	Protocol    string `json:"protocol"`
	CommunityID string `json:"community_id"`
	Bytes       int    `json:"bytes"`
}
type URL struct {
	Path   string `json:"path"`
	Query  string `json:"query"`
	Full   string `json:"full"`
	Scheme string `json:"scheme"`
	Domain string `json:"domain"`
}
type Client struct {
	Bytes int    `json:"bytes"`
	IP    string `json:"ip"`
	Port  int    `json:"port"`
}
type Event struct {
	Duration int       `json:"duration"`
	Start    time.Time `json:"start"`
	End      time.Time `json:"end"`
	Kind     string    `json:"kind"`
	Category string    `json:"category"`
	Dataset  string    `json:"dataset"`
}
type UserAgent struct {
	Original string `json:"original"`
}
type Destination struct {
	IP     string `json:"ip"`
	Port   int    `json:"port"`
	Domain string `json:"domain"`
	Bytes  int    `json:"bytes"`
}
type Agent struct {
	Hostname    string `json:"hostname"`
	ID          string `json:"id"`
	Version     string `json:"version"`
	Type        string `json:"type"`
	EphemeralID string `json:"ephemeral_id"`
}
