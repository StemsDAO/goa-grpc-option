package testdata

var MixedEndpointsConnConfigurerStructCode = `// ConnConfigurer holds the websocket connection configurer functions for the
// streaming endpoints in "StreamingResultService" service.
type ConnConfigurer struct {
	StreamingResultMethodFn goahttp.ConnConfigureFunc
}
`

var MixedEndpointsConnConfigurerInitCode = `// NewConnConfigurer initializes the websocket connection configurer function
// with fn for all the streaming endpoints in "StreamingResultService" service.
func NewConnConfigurer(fn goahttp.ConnConfigureFunc) *ConnConfigurer {
	return &ConnConfigurer{
		StreamingResultMethodFn: fn,
	}
}
`

var StreamingResultServerHandlerInitCode = `// NewStreamingResultMethodHandler creates a HTTP handler which loads the HTTP
// request and calls the "StreamingResultService" service
// "StreamingResultMethod" endpoint.
func NewStreamingResultMethodHandler(
	endpoint goa.Endpoint,
	mux goahttp.Muxer,
	decoder func(*http.Request) goahttp.Decoder,
	encoder func(context.Context, http.ResponseWriter) goahttp.Encoder,
	errhandler func(context.Context, http.ResponseWriter, error),
	formatter func(err error) goahttp.Statuser,
	upgrader goahttp.Upgrader,
	configurer goahttp.ConnConfigureFunc,
) http.Handler {
	var (
		decodeRequest = DecodeStreamingResultMethodRequest(mux, decoder)
		encodeError   = goahttp.ErrorEncoder(encoder, formatter)
	)
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := context.WithValue(r.Context(), goahttp.AcceptTypeKey, r.Header.Get("Accept"))
		ctx = context.WithValue(ctx, goa.MethodKey, "StreamingResultMethod")
		ctx = context.WithValue(ctx, goa.ServiceKey, "StreamingResultService")
		payload, err := decodeRequest(r)
		if err != nil {
			if err := encodeError(ctx, w, err); err != nil {
				errhandler(ctx, w, err)
			}
			return
		}
		var cancel context.CancelFunc
		ctx, cancel = context.WithCancel(ctx)
		v := &streamingresultservice.StreamingResultMethodEndpointInput{
			Stream: &StreamingResultMethodServerStream{
				upgrader:   upgrader,
				configurer: configurer,
				cancel:     cancel,
				w:          w,
				r:          r,
			},
			Payload: payload.(*streamingresultservice.Request),
		}
		_, err = endpoint(ctx, v)
		if err != nil {
			if _, werr := w.Write(nil); werr == http.ErrHijacked {
				// Response writer has been hijacked, do not encode the error
				errhandler(ctx, w, err)
				return
			}
			if err := encodeError(ctx, w, err); err != nil {
				errhandler(ctx, w, err)
			}
			return
		}
	})
}
`

var StreamingResultServerStreamSendCode = `// Send streams instances of "streamingresultservice.UserType" to the
// "StreamingResultMethod" endpoint websocket connection.
func (s *StreamingResultMethodServerStream) Send(v *streamingresultservice.UserType) error {
	var err error
	// Upgrade the HTTP connection to a websocket connection only once. Connection
	// upgrade is done here so that authorization logic in the endpoint is executed
	// before calling the actual service method which may call Send().
	s.once.Do(func() {
		var conn *websocket.Conn
		conn, err = s.upgrader.Upgrade(s.w, s.r, nil)
		if err != nil {
			return
		}
		if s.configurer != nil {
			conn = s.configurer(conn, s.cancel)
		}
		s.conn = conn
	})
	if err != nil {
		return err
	}
	res := v
	body := NewStreamingResultMethodResponseBody(res)
	return s.conn.WriteJSON(body)
}
`

var StreamingResultServerStreamCloseCode = `// Close closes the "StreamingResultMethod" endpoint websocket connection.
func (s *StreamingResultMethodServerStream) Close() error {
	var err error
	if s.conn == nil {
		return nil
	}
	if err = s.conn.WriteControl(
		websocket.CloseMessage,
		websocket.FormatCloseMessage(websocket.CloseNormalClosure, "server closing connection"),
		time.Now().Add(time.Second),
	); err != nil {
		return err
	}
	return s.conn.Close()
}
`

var StreamingResultWithViewsServerStreamSendCode = `// Send streams instances of "streamingresultwithviewsservice.Usertype" to the
// "StreamingResultWithViewsMethod" endpoint websocket connection.
func (s *StreamingResultWithViewsMethodServerStream) Send(v *streamingresultwithviewsservice.Usertype) error {
	var err error
	// Upgrade the HTTP connection to a websocket connection only once. Connection
	// upgrade is done here so that authorization logic in the endpoint is executed
	// before calling the actual service method which may call Send().
	s.once.Do(func() {
		respHdr := make(http.Header)
		respHdr.Add("goa-view", s.view)
		var conn *websocket.Conn
		conn, err = s.upgrader.Upgrade(s.w, s.r, respHdr)
		if err != nil {
			return
		}
		if s.configurer != nil {
			conn = s.configurer(conn, s.cancel)
		}
		s.conn = conn
	})
	if err != nil {
		return err
	}
	res := streamingresultwithviewsservice.NewViewedUsertype(v, s.view)
	var body interface{}
	switch s.view {
	case "tiny":
		body = NewStreamingResultWithViewsMethodResponseBodyTiny(res.Projected)
	case "extended":
		body = NewStreamingResultWithViewsMethodResponseBodyExtended(res.Projected)
	case "default", "":
		body = NewStreamingResultWithViewsMethodResponseBody(res.Projected)
	}
	return s.conn.WriteJSON(body)
}
`

var StreamingResultWithViewsServerStreamSetViewCode = `// SetView sets the view to render the streamingresultwithviewsservice.Usertype
// type before sending to the "StreamingResultWithViewsMethod" endpoint
// websocket connection.
func (s *StreamingResultWithViewsMethodServerStream) SetView(view string) {
	s.view = view
}
`

var StreamingResultNoPayloadServerHandlerInitCode = `// NewStreamingResultNoPayloadMethodHandler creates a HTTP handler which loads
// the HTTP request and calls the "StreamingResultNoPayloadService" service
// "StreamingResultNoPayloadMethod" endpoint.
func NewStreamingResultNoPayloadMethodHandler(
	endpoint goa.Endpoint,
	mux goahttp.Muxer,
	decoder func(*http.Request) goahttp.Decoder,
	encoder func(context.Context, http.ResponseWriter) goahttp.Encoder,
	errhandler func(context.Context, http.ResponseWriter, error),
	formatter func(err error) goahttp.Statuser,
	upgrader goahttp.Upgrader,
	configurer goahttp.ConnConfigureFunc,
) http.Handler {
	var (
		encodeError = goahttp.ErrorEncoder(encoder, formatter)
	)
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := context.WithValue(r.Context(), goahttp.AcceptTypeKey, r.Header.Get("Accept"))
		ctx = context.WithValue(ctx, goa.MethodKey, "StreamingResultNoPayloadMethod")
		ctx = context.WithValue(ctx, goa.ServiceKey, "StreamingResultNoPayloadService")
		var err error
		var cancel context.CancelFunc
		ctx, cancel = context.WithCancel(ctx)
		v := &streamingresultnopayloadservice.StreamingResultNoPayloadMethodEndpointInput{
			Stream: &StreamingResultNoPayloadMethodServerStream{
				upgrader:   upgrader,
				configurer: configurer,
				cancel:     cancel,
				w:          w,
				r:          r,
			},
		}
		_, err = endpoint(ctx, v)
		if err != nil {
			if _, werr := w.Write(nil); werr == http.ErrHijacked {
				// Response writer has been hijacked, do not encode the error
				errhandler(ctx, w, err)
				return
			}
			if err := encodeError(ctx, w, err); err != nil {
				errhandler(ctx, w, err)
			}
			return
		}
	})
}
`

var StreamingResultClientEndpointCode = `// StreamingResultMethod returns an endpoint that makes HTTP requests to the
// StreamingResultService service StreamingResultMethod server.
func (c *Client) StreamingResultMethod() goa.Endpoint {
	var (
		decodeResponse = DecodeStreamingResultMethodResponse(c.decoder, c.RestoreResponseBody)
	)
	return func(ctx context.Context, v interface{}) (interface{}, error) {
		req, err := c.BuildStreamingResultMethodRequest(ctx, v)
		if err != nil {
			return nil, err
		}
		var cancel context.CancelFunc
		ctx, cancel = context.WithCancel(ctx)
		conn, resp, err := c.dialer.DialContext(ctx, req.URL.String(), req.Header)
		if err != nil {
			if resp != nil {
				return decodeResponse(resp)
			}
			return nil, goahttp.ErrRequestError("StreamingResultService", "StreamingResultMethod", err)
		}
		if c.configurer.StreamingResultMethodFn != nil {
			conn = c.configurer.StreamingResultMethodFn(conn, cancel)
		}
		go func() {
			<-ctx.Done()
			conn.WriteControl(
				websocket.CloseMessage,
				websocket.FormatCloseMessage(websocket.CloseNormalClosure, "client closing connection"),
				time.Now().Add(time.Second),
			)
			conn.Close()
		}()
		stream := &StreamingResultMethodClientStream{conn: conn}
		return stream, nil
	}
}
`

var StreamingResultWithViewsServerStreamCloseCode = `// Close closes the "StreamingResultWithViewsMethod" endpoint websocket
// connection.
func (s *StreamingResultWithViewsMethodServerStream) Close() error {
	var err error
	if s.conn == nil {
		return nil
	}
	if err = s.conn.WriteControl(
		websocket.CloseMessage,
		websocket.FormatCloseMessage(websocket.CloseNormalClosure, "server closing connection"),
		time.Now().Add(time.Second),
	); err != nil {
		return err
	}
	return s.conn.Close()
}
`

var StreamingResultClientStreamRecvCode = `// Recv reads instances of "streamingresultservice.UserType" from the
// "StreamingResultMethod" endpoint websocket connection.
func (s *StreamingResultMethodClientStream) Recv() (*streamingresultservice.UserType, error) {
	var (
		rv   *streamingresultservice.UserType
		body StreamingResultMethodResponseBody
		err  error
	)
	err = s.conn.ReadJSON(&body)
	if websocket.IsCloseError(err, websocket.CloseNormalClosure) {
		s.conn.Close()
		return rv, io.EOF
	}
	if err != nil {
		return rv, err
	}
	res := NewStreamingResultMethodUserTypeOK(&body)
	return res, nil
}
`

var StreamingResultWithViewsClientEndpointCode = `// StreamingResultWithViewsMethod returns an endpoint that makes HTTP requests
// to the StreamingResultWithViewsService service
// StreamingResultWithViewsMethod server.
func (c *Client) StreamingResultWithViewsMethod() goa.Endpoint {
	var (
		decodeResponse = DecodeStreamingResultWithViewsMethodResponse(c.decoder, c.RestoreResponseBody)
	)
	return func(ctx context.Context, v interface{}) (interface{}, error) {
		req, err := c.BuildStreamingResultWithViewsMethodRequest(ctx, v)
		if err != nil {
			return nil, err
		}
		var cancel context.CancelFunc
		ctx, cancel = context.WithCancel(ctx)
		conn, resp, err := c.dialer.DialContext(ctx, req.URL.String(), req.Header)
		if err != nil {
			if resp != nil {
				return decodeResponse(resp)
			}
			return nil, goahttp.ErrRequestError("StreamingResultWithViewsService", "StreamingResultWithViewsMethod", err)
		}
		if c.configurer.StreamingResultWithViewsMethodFn != nil {
			conn = c.configurer.StreamingResultWithViewsMethodFn(conn, cancel)
		}
		go func() {
			<-ctx.Done()
			conn.WriteControl(
				websocket.CloseMessage,
				websocket.FormatCloseMessage(websocket.CloseNormalClosure, "client closing connection"),
				time.Now().Add(time.Second),
			)
			conn.Close()
		}()
		stream := &StreamingResultWithViewsMethodClientStream{conn: conn}
		view := resp.Header.Get("goa-view")
		stream.SetView(view)
		return stream, nil
	}
}
`

var StreamingResultWithViewsClientStreamRecvCode = `// Recv reads instances of "streamingresultwithviewsservice.Usertype" from the
// "StreamingResultWithViewsMethod" endpoint websocket connection.
func (s *StreamingResultWithViewsMethodClientStream) Recv() (*streamingresultwithviewsservice.Usertype, error) {
	var (
		rv   *streamingresultwithviewsservice.Usertype
		body StreamingResultWithViewsMethodResponseBody
		err  error
	)
	err = s.conn.ReadJSON(&body)
	if websocket.IsCloseError(err, websocket.CloseNormalClosure) {
		s.conn.Close()
		return rv, io.EOF
	}
	if err != nil {
		return rv, err
	}
	res := NewStreamingResultWithViewsMethodUsertypeOK(&body)
	vres := &streamingresultwithviewsserviceviews.Usertype{res, s.view}
	if err := streamingresultwithviewsserviceviews.ValidateUsertype(vres); err != nil {
		return rv, goahttp.ErrValidationError("StreamingResultWithViewsService", "StreamingResultWithViewsMethod", err)
	}
	return streamingresultwithviewsservice.NewUsertype(vres), nil
}
`

var StreamingResultWithViewsClientStreamSetViewCode = `// SetView sets the view to render the  type before sending to the
// "StreamingResultWithViewsMethod" endpoint websocket connection.
func (s *StreamingResultWithViewsMethodClientStream) SetView(view string) {
	s.view = view
}
`

var StreamingResultWithExplicitViewClientEndpointCode = `// StreamingResultWithExplicitViewMethod returns an endpoint that makes HTTP
// requests to the StreamingResultWithExplicitViewService service
// StreamingResultWithExplicitViewMethod server.
func (c *Client) StreamingResultWithExplicitViewMethod() goa.Endpoint {
	var (
		decodeResponse = DecodeStreamingResultWithExplicitViewMethodResponse(c.decoder, c.RestoreResponseBody)
	)
	return func(ctx context.Context, v interface{}) (interface{}, error) {
		req, err := c.BuildStreamingResultWithExplicitViewMethodRequest(ctx, v)
		if err != nil {
			return nil, err
		}
		var cancel context.CancelFunc
		ctx, cancel = context.WithCancel(ctx)
		conn, resp, err := c.dialer.DialContext(ctx, req.URL.String(), req.Header)
		if err != nil {
			if resp != nil {
				return decodeResponse(resp)
			}
			return nil, goahttp.ErrRequestError("StreamingResultWithExplicitViewService", "StreamingResultWithExplicitViewMethod", err)
		}
		if c.configurer.StreamingResultWithExplicitViewMethodFn != nil {
			conn = c.configurer.StreamingResultWithExplicitViewMethodFn(conn, cancel)
		}
		go func() {
			<-ctx.Done()
			conn.WriteControl(
				websocket.CloseMessage,
				websocket.FormatCloseMessage(websocket.CloseNormalClosure, "client closing connection"),
				time.Now().Add(time.Second),
			)
			conn.Close()
		}()
		stream := &StreamingResultWithExplicitViewMethodClientStream{conn: conn}
		return stream, nil
	}
}
`

var StreamingResultWithExplicitViewClientStreamRecvCode = `// Recv reads instances of "streamingresultwithexplicitviewservice.Usertype"
// from the "StreamingResultWithExplicitViewMethod" endpoint websocket
// connection.
func (s *StreamingResultWithExplicitViewMethodClientStream) Recv() (*streamingresultwithexplicitviewservice.Usertype, error) {
	var (
		rv   *streamingresultwithexplicitviewservice.Usertype
		body StreamingResultWithExplicitViewMethodResponseBody
		err  error
	)
	err = s.conn.ReadJSON(&body)
	if websocket.IsCloseError(err, websocket.CloseNormalClosure) {
		s.conn.Close()
		return rv, io.EOF
	}
	if err != nil {
		return rv, err
	}
	res := NewStreamingResultWithExplicitViewMethodUsertypeOK(&body)
	vres := &streamingresultwithexplicitviewserviceviews.Usertype{res, "extended"}
	if err := streamingresultwithexplicitviewserviceviews.ValidateUsertype(vres); err != nil {
		return rv, goahttp.ErrValidationError("StreamingResultWithExplicitViewService", "StreamingResultWithExplicitViewMethod", err)
	}
	return streamingresultwithexplicitviewservice.NewUsertype(vres), nil
}
`

var StreamingResultWithExplicitViewServerStreamSendCode = `// Send streams instances of "streamingresultwithexplicitviewservice.Usertype"
// to the "StreamingResultWithExplicitViewMethod" endpoint websocket connection.
func (s *StreamingResultWithExplicitViewMethodServerStream) Send(v *streamingresultwithexplicitviewservice.Usertype) error {
	var err error
	// Upgrade the HTTP connection to a websocket connection only once. Connection
	// upgrade is done here so that authorization logic in the endpoint is executed
	// before calling the actual service method which may call Send().
	s.once.Do(func() {
		var conn *websocket.Conn
		conn, err = s.upgrader.Upgrade(s.w, s.r, nil)
		if err != nil {
			return
		}
		if s.configurer != nil {
			conn = s.configurer(conn, s.cancel)
		}
		s.conn = conn
	})
	if err != nil {
		return err
	}
	res := streamingresultwithexplicitviewservice.NewViewedUsertype(v, "extended")
	body := NewStreamingResultWithExplicitViewMethodResponseBodyExtended(res.Projected)
	return s.conn.WriteJSON(body)
}
`

var StreamingResultListWithViewsServerStreamSendCode = `// Send streams instances of
// "streamingresultlistwithviewsservice.UsertypeList" to the
// "StreamingResultListWithViewsMethod" endpoint websocket connection.
func (s *StreamingResultListWithViewsMethodServerStream) Send(v streamingresultlistwithviewsservice.UsertypeList) error {
	var err error
	// Upgrade the HTTP connection to a websocket connection only once. Connection
	// upgrade is done here so that authorization logic in the endpoint is executed
	// before calling the actual service method which may call Send().
	s.once.Do(func() {
		respHdr := make(http.Header)
		respHdr.Add("goa-view", s.view)
		var conn *websocket.Conn
		conn, err = s.upgrader.Upgrade(s.w, s.r, respHdr)
		if err != nil {
			return
		}
		if s.configurer != nil {
			conn = s.configurer(conn, s.cancel)
		}
		s.conn = conn
	})
	if err != nil {
		return err
	}
	res := streamingresultlistwithviewsservice.NewViewedUsertypeList(v, s.view)
	var body interface{}
	switch s.view {
	case "tiny":
		body = NewUsertypeResponseTinyList(res.Projected)
	case "extended":
		body = NewUsertypeResponseExtendedList(res.Projected)
	case "default", "":
		body = NewUsertypeResponseList(res.Projected)
	}
	return s.conn.WriteJSON(body)
}
`

var StreamingResultListWithViewsServerStreamSetViewCode = `// SetView sets the view to render the
// streamingresultlistwithviewsservice.UsertypeList type before
// sending to the "StreamingResultListWithViewsMethod" endpoint websocket
// connection.
func (s *StreamingResultListWithViewsMethodServerStream) SetView(view string) {
	s.view = view
}
`

var StreamingResultListWithViewsClientStreamRecvCode = `// Recv reads instances of
// "streamingresultlistwithviewsservice.UsertypeList" from the
// "StreamingResultListWithViewsMethod" endpoint websocket connection.
func (s *StreamingResultListWithViewsMethodClientStream) Recv() (streamingresultlistwithviewsservice.UsertypeList, error) {
	var (
		rv   streamingresultlistwithviewsservice.UsertypeList
		body StreamingResultListWithViewsMethodResponseBody
		err  error
	)
	err = s.conn.ReadJSON(&body)
	if websocket.IsCloseError(err, websocket.CloseNormalClosure) {
		s.conn.Close()
		return rv, io.EOF
	}
	if err != nil {
		return rv, err
	}
	res := NewStreamingResultListWithViewsMethodUsertypeListOK(body)
	vres := streamingresultlistwithviewsserviceviews.UsertypeList{res, s.view}
	if err := streamingresultlistwithviewsserviceviews.ValidateUsertypeList(vres); err != nil {
		return rv, goahttp.ErrValidationError("StreamingResultListWithViewsService", "StreamingResultListWithViewsMethod", err)
	}
	return streamingresultlistwithviewsservice.NewUsertypeList(vres), nil
}
`

var StreamingResultListWithViewsClientStreamSetViewCode = `// SetView sets the view to render the  type before sending to the
// "StreamingResultListWithViewsMethod" endpoint websocket connection.
func (s *StreamingResultListWithViewsMethodClientStream) SetView(view string) {
	s.view = view
}
`

var StreamingResultListWithExplicitViewServerStreamSendCode = `// Send streams instances of
// "streamingresultlistwithexplicitviewservice.UsertypeList" to the
// "StreamingResultListWithExplicitViewMethod" endpoint websocket
// connection.
func (s *StreamingResultListWithExplicitViewMethodServerStream) Send(v streamingresultlistwithexplicitviewservice.UsertypeList) error {
	var err error
	// Upgrade the HTTP connection to a websocket connection only once. Connection
	// upgrade is done here so that authorization logic in the endpoint is executed
	// before calling the actual service method which may call Send().
	s.once.Do(func() {
		var conn *websocket.Conn
		conn, err = s.upgrader.Upgrade(s.w, s.r, nil)
		if err != nil {
			return
		}
		if s.configurer != nil {
			conn = s.configurer(conn, s.cancel)
		}
		s.conn = conn
	})
	if err != nil {
		return err
	}
	res := streamingresultlistwithexplicitviewservice.NewViewedUsertypeList(v, "tiny")
	body := NewUsertypeResponseTinyList(res.Projected)
	return s.conn.WriteJSON(body)
}
`

var StreamingResultListWithExplicitViewClientEndpointCode = `// StreamingResultListWithExplicitViewMethod returns an endpoint that
// makes HTTP requests to the StreamingResultListWithExplicitViewService
// service StreamingResultListWithExplicitViewMethod server.
func (c *Client) StreamingResultListWithExplicitViewMethod() goa.Endpoint {
	var (
		decodeResponse = DecodeStreamingResultListWithExplicitViewMethodResponse(c.decoder, c.RestoreResponseBody)
	)
	return func(ctx context.Context, v interface{}) (interface{}, error) {
		req, err := c.BuildStreamingResultListWithExplicitViewMethodRequest(ctx, v)
		if err != nil {
			return nil, err
		}
		var cancel context.CancelFunc
		ctx, cancel = context.WithCancel(ctx)
		conn, resp, err := c.dialer.DialContext(ctx, req.URL.String(), req.Header)
		if err != nil {
			if resp != nil {
				return decodeResponse(resp)
			}
			return nil, goahttp.ErrRequestError("StreamingResultListWithExplicitViewService", "StreamingResultListWithExplicitViewMethod", err)
		}
		if c.configurer.StreamingResultListWithExplicitViewMethodFn != nil {
			conn = c.configurer.StreamingResultListWithExplicitViewMethodFn(conn, cancel)
		}
		go func() {
			<-ctx.Done()
			conn.WriteControl(
				websocket.CloseMessage,
				websocket.FormatCloseMessage(websocket.CloseNormalClosure, "client closing connection"),
				time.Now().Add(time.Second),
			)
			conn.Close()
		}()
		stream := &StreamingResultListWithExplicitViewMethodClientStream{conn: conn}
		return stream, nil
	}
}
`

var StreamingResultListWithExplicitViewClientStreamRecvCode = `// Recv reads instances of
// "streamingresultlistwithexplicitviewservice.UsertypeList" from
// the "StreamingResultListWithExplicitViewMethod" endpoint websocket
// connection.
func (s *StreamingResultListWithExplicitViewMethodClientStream) Recv() (streamingresultlistwithexplicitviewservice.UsertypeList, error) {
	var (
		rv   streamingresultlistwithexplicitviewservice.UsertypeList
		body StreamingResultListWithExplicitViewMethodResponseBody
		err  error
	)
	err = s.conn.ReadJSON(&body)
	if websocket.IsCloseError(err, websocket.CloseNormalClosure) {
		s.conn.Close()
		return rv, io.EOF
	}
	if err != nil {
		return rv, err
	}
	res := NewStreamingResultListWithExplicitViewMethodUsertypeListOK(body)
	vres := streamingresultlistwithexplicitviewserviceviews.UsertypeList{res, "tiny"}
	if err := streamingresultlistwithexplicitviewserviceviews.ValidateUsertypeList(vres); err != nil {
		return rv, goahttp.ErrValidationError("StreamingResultListWithExplicitViewService", "StreamingResultListWithExplicitViewMethod", err)
	}
	return streamingresultlistwithexplicitviewservice.NewUsertypeList(vres), nil
}
`

var StreamingResultPrimitiveServerStreamSendCode = `// Send streams instances of "string" to the "StreamingResultPrimitiveMethod"
// endpoint websocket connection.
func (s *StreamingResultPrimitiveMethodServerStream) Send(v string) error {
	var err error
	// Upgrade the HTTP connection to a websocket connection only once. Connection
	// upgrade is done here so that authorization logic in the endpoint is executed
	// before calling the actual service method which may call Send().
	s.once.Do(func() {
		var conn *websocket.Conn
		conn, err = s.upgrader.Upgrade(s.w, s.r, nil)
		if err != nil {
			return
		}
		if s.configurer != nil {
			conn = s.configurer(conn, s.cancel)
		}
		s.conn = conn
	})
	if err != nil {
		return err
	}
	res := v
	return s.conn.WriteJSON(res)
}
`

var StreamingResultPrimitiveClientStreamRecvCode = `// Recv reads instances of "string" from the "StreamingResultPrimitiveMethod"
// endpoint websocket connection.
func (s *StreamingResultPrimitiveMethodClientStream) Recv() (string, error) {
	var (
		rv   string
		body string
		err  error
	)
	err = s.conn.ReadJSON(&body)
	if websocket.IsCloseError(err, websocket.CloseNormalClosure) {
		s.conn.Close()
		return rv, io.EOF
	}
	if err != nil {
		return rv, err
	}
	return body, nil
}
`

var StreamingResultPrimitiveArrayServerStreamSendCode = `// Send streams instances of "[]int32" to the
// "StreamingResultPrimitiveArrayMethod" endpoint websocket connection.
func (s *StreamingResultPrimitiveArrayMethodServerStream) Send(v []int32) error {
	var err error
	// Upgrade the HTTP connection to a websocket connection only once. Connection
	// upgrade is done here so that authorization logic in the endpoint is executed
	// before calling the actual service method which may call Send().
	s.once.Do(func() {
		var conn *websocket.Conn
		conn, err = s.upgrader.Upgrade(s.w, s.r, nil)
		if err != nil {
			return
		}
		if s.configurer != nil {
			conn = s.configurer(conn, s.cancel)
		}
		s.conn = conn
	})
	if err != nil {
		return err
	}
	res := v
	return s.conn.WriteJSON(res)
}
`

var StreamingResultPrimitiveArrayClientStreamRecvCode = `// Recv reads instances of "[]int32" from the
// "StreamingResultPrimitiveArrayMethod" endpoint websocket connection.
func (s *StreamingResultPrimitiveArrayMethodClientStream) Recv() ([]int32, error) {
	var (
		rv   []int32
		body []int32
		err  error
	)
	err = s.conn.ReadJSON(&body)
	if websocket.IsCloseError(err, websocket.CloseNormalClosure) {
		s.conn.Close()
		return rv, io.EOF
	}
	if err != nil {
		return rv, err
	}
	return body, nil
}
`

var StreamingResultPrimitiveMapServerStreamSendCode = `// Send streams instances of "map[int32]string" to the
// "StreamingResultPrimitiveMapMethod" endpoint websocket connection.
func (s *StreamingResultPrimitiveMapMethodServerStream) Send(v map[int32]string) error {
	var err error
	// Upgrade the HTTP connection to a websocket connection only once. Connection
	// upgrade is done here so that authorization logic in the endpoint is executed
	// before calling the actual service method which may call Send().
	s.once.Do(func() {
		var conn *websocket.Conn
		conn, err = s.upgrader.Upgrade(s.w, s.r, nil)
		if err != nil {
			return
		}
		if s.configurer != nil {
			conn = s.configurer(conn, s.cancel)
		}
		s.conn = conn
	})
	if err != nil {
		return err
	}
	res := v
	return s.conn.WriteJSON(res)
}
`

var StreamingResultPrimitiveMapClientStreamRecvCode = `// Recv reads instances of "map[int32]string" from the
// "StreamingResultPrimitiveMapMethod" endpoint websocket connection.
func (s *StreamingResultPrimitiveMapMethodClientStream) Recv() (map[int32]string, error) {
	var (
		rv   map[int32]string
		body map[int32]string
		err  error
	)
	err = s.conn.ReadJSON(&body)
	if websocket.IsCloseError(err, websocket.CloseNormalClosure) {
		s.conn.Close()
		return rv, io.EOF
	}
	if err != nil {
		return rv, err
	}
	return body, nil
}
`

var StreamingResultUserTypeArrayServerStreamSendCode = `// Send streams instances of "[]*streamingresultusertypearrayservice.UserType"
// to the "StreamingResultUserTypeArrayMethod" endpoint websocket connection.
func (s *StreamingResultUserTypeArrayMethodServerStream) Send(v []*streamingresultusertypearrayservice.UserType) error {
	var err error
	// Upgrade the HTTP connection to a websocket connection only once. Connection
	// upgrade is done here so that authorization logic in the endpoint is executed
	// before calling the actual service method which may call Send().
	s.once.Do(func() {
		var conn *websocket.Conn
		conn, err = s.upgrader.Upgrade(s.w, s.r, nil)
		if err != nil {
			return
		}
		if s.configurer != nil {
			conn = s.configurer(conn, s.cancel)
		}
		s.conn = conn
	})
	if err != nil {
		return err
	}
	res := v
	body := NewStreamingResultUserTypeArrayMethodResponseBody(res)
	return s.conn.WriteJSON(body)
}
`

var StreamingResultUserTypeArrayClientStreamRecvCode = `// Recv reads instances of "[]*streamingresultusertypearrayservice.UserType"
// from the "StreamingResultUserTypeArrayMethod" endpoint websocket connection.
func (s *StreamingResultUserTypeArrayMethodClientStream) Recv() ([]*streamingresultusertypearrayservice.UserType, error) {
	var (
		rv   []*streamingresultusertypearrayservice.UserType
		body StreamingResultUserTypeArrayMethodResponseBody
		err  error
	)
	err = s.conn.ReadJSON(&body)
	if websocket.IsCloseError(err, websocket.CloseNormalClosure) {
		s.conn.Close()
		return rv, io.EOF
	}
	if err != nil {
		return rv, err
	}
	res := NewStreamingResultUserTypeArrayMethodUserTypeOK(body)
	return res, nil
}
`

var StreamingResultUserTypeMapServerStreamSendCode = `// Send streams instances of
// "map[string]*streamingresultusertypemapservice.UserType" to the
// "StreamingResultUserTypeMapMethod" endpoint websocket connection.
func (s *StreamingResultUserTypeMapMethodServerStream) Send(v map[string]*streamingresultusertypemapservice.UserType) error {
	var err error
	// Upgrade the HTTP connection to a websocket connection only once. Connection
	// upgrade is done here so that authorization logic in the endpoint is executed
	// before calling the actual service method which may call Send().
	s.once.Do(func() {
		var conn *websocket.Conn
		conn, err = s.upgrader.Upgrade(s.w, s.r, nil)
		if err != nil {
			return
		}
		if s.configurer != nil {
			conn = s.configurer(conn, s.cancel)
		}
		s.conn = conn
	})
	if err != nil {
		return err
	}
	res := v
	body := NewStreamingResultUserTypeMapMethodResponseBody(res)
	return s.conn.WriteJSON(body)
}
`

var StreamingResultUserTypeMapClientStreamRecvCode = `// Recv reads instances of
// "map[string]*streamingresultusertypemapservice.UserType" from the
// "StreamingResultUserTypeMapMethod" endpoint websocket connection.
func (s *StreamingResultUserTypeMapMethodClientStream) Recv() (map[string]*streamingresultusertypemapservice.UserType, error) {
	var (
		rv   map[string]*streamingresultusertypemapservice.UserType
		body StreamingResultUserTypeMapMethodResponseBody
		err  error
	)
	err = s.conn.ReadJSON(&body)
	if websocket.IsCloseError(err, websocket.CloseNormalClosure) {
		s.conn.Close()
		return rv, io.EOF
	}
	if err != nil {
		return rv, err
	}
	res := NewStreamingResultUserTypeMapMethodMapStringUserTypeOK(body)
	return res, nil
}
`

var StreamingResultNoPayloadClientEndpointCode = `// StreamingResultNoPayloadMethod returns an endpoint that makes HTTP requests
// to the StreamingResultNoPayloadService service
// StreamingResultNoPayloadMethod server.
func (c *Client) StreamingResultNoPayloadMethod() goa.Endpoint {
	var (
		decodeResponse = DecodeStreamingResultNoPayloadMethodResponse(c.decoder, c.RestoreResponseBody)
	)
	return func(ctx context.Context, v interface{}) (interface{}, error) {
		req, err := c.BuildStreamingResultNoPayloadMethodRequest(ctx, v)
		if err != nil {
			return nil, err
		}
		var cancel context.CancelFunc
		ctx, cancel = context.WithCancel(ctx)
		conn, resp, err := c.dialer.DialContext(ctx, req.URL.String(), req.Header)
		if err != nil {
			if resp != nil {
				return decodeResponse(resp)
			}
			return nil, goahttp.ErrRequestError("StreamingResultNoPayloadService", "StreamingResultNoPayloadMethod", err)
		}
		if c.configurer.StreamingResultNoPayloadMethodFn != nil {
			conn = c.configurer.StreamingResultNoPayloadMethodFn(conn, cancel)
		}
		go func() {
			<-ctx.Done()
			conn.WriteControl(
				websocket.CloseMessage,
				websocket.FormatCloseMessage(websocket.CloseNormalClosure, "client closing connection"),
				time.Now().Add(time.Second),
			)
			conn.Close()
		}()
		stream := &StreamingResultNoPayloadMethodClientStream{conn: conn}
		return stream, nil
	}
}
`

var StreamingPayloadServerHandlerInitCode = `// NewStreamingPayloadMethodHandler creates a HTTP handler which loads the HTTP
// request and calls the "StreamingPayloadService" service
// "StreamingPayloadMethod" endpoint.
func NewStreamingPayloadMethodHandler(
	endpoint goa.Endpoint,
	mux goahttp.Muxer,
	decoder func(*http.Request) goahttp.Decoder,
	encoder func(context.Context, http.ResponseWriter) goahttp.Encoder,
	errhandler func(context.Context, http.ResponseWriter, error),
	formatter func(err error) goahttp.Statuser,
	upgrader goahttp.Upgrader,
	configurer goahttp.ConnConfigureFunc,
) http.Handler {
	var (
		decodeRequest = DecodeStreamingPayloadMethodRequest(mux, decoder)
		encodeError   = goahttp.ErrorEncoder(encoder, formatter)
	)
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := context.WithValue(r.Context(), goahttp.AcceptTypeKey, r.Header.Get("Accept"))
		ctx = context.WithValue(ctx, goa.MethodKey, "StreamingPayloadMethod")
		ctx = context.WithValue(ctx, goa.ServiceKey, "StreamingPayloadService")
		payload, err := decodeRequest(r)
		if err != nil {
			if err := encodeError(ctx, w, err); err != nil {
				errhandler(ctx, w, err)
			}
			return
		}
		var cancel context.CancelFunc
		ctx, cancel = context.WithCancel(ctx)
		v := &streamingpayloadservice.StreamingPayloadMethodEndpointInput{
			Stream: &StreamingPayloadMethodServerStream{
				upgrader:   upgrader,
				configurer: configurer,
				cancel:     cancel,
				w:          w,
				r:          r,
			},
			Payload: payload.(*streamingpayloadservice.Payload),
		}
		_, err = endpoint(ctx, v)
		if err != nil {
			if _, werr := w.Write(nil); werr == http.ErrHijacked {
				// Response writer has been hijacked, do not encode the error
				errhandler(ctx, w, err)
				return
			}
			if err := encodeError(ctx, w, err); err != nil {
				errhandler(ctx, w, err)
			}
			return
		}
	})
}
`

var StreamingPayloadServerStreamSendCode = `// SendAndClose streams instances of "streamingpayloadservice.UserType" to the
// "StreamingPayloadMethod" endpoint websocket connection and closes the
// connection.
func (s *StreamingPayloadMethodServerStream) SendAndClose(v *streamingpayloadservice.UserType) error {
	defer s.conn.Close()
	res := v
	body := NewStreamingPayloadMethodResponseBody(res)
	return s.conn.WriteJSON(body)
}
`

var StreamingPayloadServerStreamRecvCode = `// Recv reads instances of "streamingpayloadservice.Request" from the
// "StreamingPayloadMethod" endpoint websocket connection.
func (s *StreamingPayloadMethodServerStream) Recv() (*streamingpayloadservice.Request, error) {
	var (
		rv  *streamingpayloadservice.Request
		msg *StreamingPayloadMethodStreamingBody
		err error
	)
	// Upgrade the HTTP connection to a websocket connection only once. Connection
	// upgrade is done here so that authorization logic in the endpoint is executed
	// before calling the actual service method which may call Recv().
	s.once.Do(func() {
		var conn *websocket.Conn
		conn, err = s.upgrader.Upgrade(s.w, s.r, nil)
		if err != nil {
			return
		}
		if s.configurer != nil {
			conn = s.configurer(conn, s.cancel)
		}
		s.conn = conn
	})
	if err != nil {
		return rv, err
	}
	if err = s.conn.ReadJSON(&msg); err != nil {
		return rv, err
	}
	if msg == nil {
		return rv, io.EOF
	}
	return NewStreamingPayloadMethodStreamingBody(msg), nil
}
`

var StreamingPayloadClientEndpointCode = `// StreamingPayloadMethod returns an endpoint that makes HTTP requests to the
// StreamingPayloadService service StreamingPayloadMethod server.
func (c *Client) StreamingPayloadMethod() goa.Endpoint {
	var (
		encodeRequest  = EncodeStreamingPayloadMethodRequest(c.encoder)
		decodeResponse = DecodeStreamingPayloadMethodResponse(c.decoder, c.RestoreResponseBody)
	)
	return func(ctx context.Context, v interface{}) (interface{}, error) {
		req, err := c.BuildStreamingPayloadMethodRequest(ctx, v)
		if err != nil {
			return nil, err
		}
		err = encodeRequest(req, v)
		if err != nil {
			return nil, err
		}
		var cancel context.CancelFunc
		ctx, cancel = context.WithCancel(ctx)
		conn, resp, err := c.dialer.DialContext(ctx, req.URL.String(), req.Header)
		if err != nil {
			if resp != nil {
				return decodeResponse(resp)
			}
			return nil, goahttp.ErrRequestError("StreamingPayloadService", "StreamingPayloadMethod", err)
		}
		if c.configurer.StreamingPayloadMethodFn != nil {
			conn = c.configurer.StreamingPayloadMethodFn(conn, cancel)
		}
		stream := &StreamingPayloadMethodClientStream{conn: conn}
		return stream, nil
	}
}
`

var StreamingPayloadClientStreamSendCode = `// Send streams instances of "streamingpayloadservice.Request" to the
// "StreamingPayloadMethod" endpoint websocket connection.
func (s *StreamingPayloadMethodClientStream) Send(v *streamingpayloadservice.Request) error {
	body := NewStreamingPayloadMethodStreamingBody(v)
	return s.conn.WriteJSON(body)
}
`

var StreamingPayloadClientStreamRecvCode = `// CloseAndRecv stops sending messages to the "StreamingPayloadMethod" endpoint
// websocket connection and reads instances of
// "streamingpayloadservice.UserType" from the connection.
func (s *StreamingPayloadMethodClientStream) CloseAndRecv() (*streamingpayloadservice.UserType, error) {
	var (
		rv   *streamingpayloadservice.UserType
		body StreamingPayloadMethodResponseBody
		err  error
	)
	defer s.conn.Close()
	// Send a nil payload to the server implying end of message
	if err = s.conn.WriteJSON(nil); err != nil {
		return rv, err
	}
	err = s.conn.ReadJSON(&body)
	if websocket.IsCloseError(err, websocket.CloseNormalClosure) {
		s.conn.Close()
		return rv, io.EOF
	}
	if err != nil {
		return rv, err
	}
	res := NewStreamingPayloadMethodUserTypeOK(&body)
	return res, nil
}
`

var StreamingPayloadNoPayloadServerHandlerInitCode = `// NewStreamingPayloadNoPayloadMethodHandler creates a HTTP handler which loads
// the HTTP request and calls the "StreamingPayloadNoPayloadService" service
// "StreamingPayloadNoPayloadMethod" endpoint.
func NewStreamingPayloadNoPayloadMethodHandler(
	endpoint goa.Endpoint,
	mux goahttp.Muxer,
	decoder func(*http.Request) goahttp.Decoder,
	encoder func(context.Context, http.ResponseWriter) goahttp.Encoder,
	errhandler func(context.Context, http.ResponseWriter, error),
	formatter func(err error) goahttp.Statuser,
	upgrader goahttp.Upgrader,
	configurer goahttp.ConnConfigureFunc,
) http.Handler {
	var (
		encodeError = goahttp.ErrorEncoder(encoder, formatter)
	)
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := context.WithValue(r.Context(), goahttp.AcceptTypeKey, r.Header.Get("Accept"))
		ctx = context.WithValue(ctx, goa.MethodKey, "StreamingPayloadNoPayloadMethod")
		ctx = context.WithValue(ctx, goa.ServiceKey, "StreamingPayloadNoPayloadService")
		var err error
		var cancel context.CancelFunc
		ctx, cancel = context.WithCancel(ctx)
		v := &streamingpayloadnopayloadservice.StreamingPayloadNoPayloadMethodEndpointInput{
			Stream: &StreamingPayloadNoPayloadMethodServerStream{
				upgrader:   upgrader,
				configurer: configurer,
				cancel:     cancel,
				w:          w,
				r:          r,
			},
		}
		_, err = endpoint(ctx, v)
		if err != nil {
			if _, werr := w.Write(nil); werr == http.ErrHijacked {
				// Response writer has been hijacked, do not encode the error
				errhandler(ctx, w, err)
				return
			}
			if err := encodeError(ctx, w, err); err != nil {
				errhandler(ctx, w, err)
			}
			return
		}
	})
}
`

var StreamingPayloadNoPayloadClientEndpointCode = `// StreamingPayloadNoPayloadMethod returns an endpoint that makes HTTP requests
// to the StreamingPayloadNoPayloadService service
// StreamingPayloadNoPayloadMethod server.
func (c *Client) StreamingPayloadNoPayloadMethod() goa.Endpoint {
	var (
		decodeResponse = DecodeStreamingPayloadNoPayloadMethodResponse(c.decoder, c.RestoreResponseBody)
	)
	return func(ctx context.Context, v interface{}) (interface{}, error) {
		req, err := c.BuildStreamingPayloadNoPayloadMethodRequest(ctx, v)
		if err != nil {
			return nil, err
		}
		var cancel context.CancelFunc
		ctx, cancel = context.WithCancel(ctx)
		conn, resp, err := c.dialer.DialContext(ctx, req.URL.String(), req.Header)
		if err != nil {
			if resp != nil {
				return decodeResponse(resp)
			}
			return nil, goahttp.ErrRequestError("StreamingPayloadNoPayloadService", "StreamingPayloadNoPayloadMethod", err)
		}
		if c.configurer.StreamingPayloadNoPayloadMethodFn != nil {
			conn = c.configurer.StreamingPayloadNoPayloadMethodFn(conn, cancel)
		}
		stream := &StreamingPayloadNoPayloadMethodClientStream{conn: conn}
		return stream, nil
	}
}
`

var StreamingPayloadNoPayloadClientStreamSendCode = `// Send streams instances of "streamingpayloadnopayloadservice.Request" to the
// "StreamingPayloadNoPayloadMethod" endpoint websocket connection.
func (s *StreamingPayloadNoPayloadMethodClientStream) Send(v *streamingpayloadnopayloadservice.Request) error {
	body := NewStreamingPayloadNoPayloadMethodStreamingBody(v)
	return s.conn.WriteJSON(body)
}
`

var StreamingPayloadNoPayloadClientStreamRecvCode = `// CloseAndRecv stops sending messages to the "StreamingPayloadNoPayloadMethod"
// endpoint websocket connection and reads instances of
// "streamingpayloadnopayloadservice.UserType" from the connection.
func (s *StreamingPayloadNoPayloadMethodClientStream) CloseAndRecv() (*streamingpayloadnopayloadservice.UserType, error) {
	var (
		rv   *streamingpayloadnopayloadservice.UserType
		body StreamingPayloadNoPayloadMethodResponseBody
		err  error
	)
	defer s.conn.Close()
	// Send a nil payload to the server implying end of message
	if err = s.conn.WriteJSON(nil); err != nil {
		return rv, err
	}
	err = s.conn.ReadJSON(&body)
	if websocket.IsCloseError(err, websocket.CloseNormalClosure) {
		s.conn.Close()
		return rv, io.EOF
	}
	if err != nil {
		return rv, err
	}
	res := NewStreamingPayloadNoPayloadMethodUserTypeOK(&body)
	return res, nil
}
`

var StreamingPayloadNoResultServerStreamRecvCode = `// Recv reads instances of "string" from the "StreamingPayloadNoResultMethod"
// endpoint websocket connection.
func (s *StreamingPayloadNoResultMethodServerStream) Recv() (string, error) {
	var (
		rv  string
		msg *string
		err error
	)
	// Upgrade the HTTP connection to a websocket connection only once. Connection
	// upgrade is done here so that authorization logic in the endpoint is executed
	// before calling the actual service method which may call Recv().
	s.once.Do(func() {
		var conn *websocket.Conn
		conn, err = s.upgrader.Upgrade(s.w, s.r, nil)
		if err != nil {
			return
		}
		if s.configurer != nil {
			conn = s.configurer(conn, s.cancel)
		}
		s.conn = conn
	})
	if err != nil {
		return rv, err
	}
	if err = s.conn.ReadJSON(&msg); err != nil {
		return rv, err
	}
	if msg == nil {
		return rv, io.EOF
	}
	return *msg, nil
}
`

var StreamingPayloadNoResultServerStreamCloseCode = `// Close closes the "StreamingPayloadNoResultMethod" endpoint websocket
// connection.
func (s *StreamingPayloadNoResultMethodServerStream) Close() error {
	var err error
	if s.conn == nil {
		return nil
	}
	if err = s.conn.WriteControl(
		websocket.CloseMessage,
		websocket.FormatCloseMessage(websocket.CloseNormalClosure, "server closing connection"),
		time.Now().Add(time.Second),
	); err != nil {
		return err
	}
	return s.conn.Close()
}
`

var StreamingPayloadNoResultClientStreamSendCode = `// Send streams instances of "string" to the "StreamingPayloadNoResultMethod"
// endpoint websocket connection.
func (s *StreamingPayloadNoResultMethodClientStream) Send(v string) error {
	return s.conn.WriteJSON(v)
}
`

var StreamingPayloadNoResultClientStreamCloseCode = `// Close closes the "StreamingPayloadNoResultMethod" endpoint websocket
// connection.
func (s *StreamingPayloadNoResultMethodClientStream) Close() error {
	var err error
	// Send a nil payload to the server implying client closing connection.
	if err = s.conn.WriteJSON(nil); err != nil {
		return err
	}
	return s.conn.Close()
}
`

var StreamingPayloadResultWithViewsServerStreamSendCode = `// SendAndClose streams instances of
// "streamingpayloadresultwithviewsservice.Usertype" to the
// "StreamingPayloadResultWithViewsMethod" endpoint websocket connection and
// closes the connection.
func (s *StreamingPayloadResultWithViewsMethodServerStream) SendAndClose(v *streamingpayloadresultwithviewsservice.Usertype) error {
	defer s.conn.Close()
	res := streamingpayloadresultwithviewsservice.NewViewedUsertype(v, s.view)
	var body interface{}
	switch s.view {
	case "tiny":
		body = NewStreamingPayloadResultWithViewsMethodResponseBodyTiny(res.Projected)
	case "extended":
		body = NewStreamingPayloadResultWithViewsMethodResponseBodyExtended(res.Projected)
	case "default", "":
		body = NewStreamingPayloadResultWithViewsMethodResponseBody(res.Projected)
	}
	return s.conn.WriteJSON(body)
}
`

var StreamingPayloadResultWithViewsServerStreamRecvCode = `// Recv reads instances of "float32" from the
// "StreamingPayloadResultWithViewsMethod" endpoint websocket connection.
func (s *StreamingPayloadResultWithViewsMethodServerStream) Recv() (float32, error) {
	var (
		rv  float32
		msg *float32
		err error
	)
	// Upgrade the HTTP connection to a websocket connection only once. Connection
	// upgrade is done here so that authorization logic in the endpoint is executed
	// before calling the actual service method which may call Recv().
	s.once.Do(func() {
		var conn *websocket.Conn
		conn, err = s.upgrader.Upgrade(s.w, s.r, nil)
		if err != nil {
			return
		}
		if s.configurer != nil {
			conn = s.configurer(conn, s.cancel)
		}
		s.conn = conn
	})
	if err != nil {
		return rv, err
	}
	if err = s.conn.ReadJSON(&msg); err != nil {
		return rv, err
	}
	if msg == nil {
		return rv, io.EOF
	}
	return *msg, nil
}
`

var StreamingPayloadResultWithViewsServerStreamSetViewCode = `// SetView sets the view to render the
// streamingpayloadresultwithviewsservice.Usertype type before sending to the
// "StreamingPayloadResultWithViewsMethod" endpoint websocket connection.
func (s *StreamingPayloadResultWithViewsMethodServerStream) SetView(view string) {
	s.view = view
}
`

var StreamingPayloadResultWithViewsClientStreamSendCode = `// Send streams instances of "float32" to the
// "StreamingPayloadResultWithViewsMethod" endpoint websocket connection.
func (s *StreamingPayloadResultWithViewsMethodClientStream) Send(v float32) error {
	return s.conn.WriteJSON(v)
}
`

var StreamingPayloadResultWithViewsClientStreamRecvCode = `// CloseAndRecv stops sending messages to the
// "StreamingPayloadResultWithViewsMethod" endpoint websocket connection and
// reads instances of "streamingpayloadresultwithviewsservice.Usertype" from
// the connection.
func (s *StreamingPayloadResultWithViewsMethodClientStream) CloseAndRecv() (*streamingpayloadresultwithviewsservice.Usertype, error) {
	var (
		rv   *streamingpayloadresultwithviewsservice.Usertype
		body StreamingPayloadResultWithViewsMethodResponseBody
		err  error
	)
	defer s.conn.Close()
	// Send a nil payload to the server implying end of message
	if err = s.conn.WriteJSON(nil); err != nil {
		return rv, err
	}
	err = s.conn.ReadJSON(&body)
	if websocket.IsCloseError(err, websocket.CloseNormalClosure) {
		s.conn.Close()
		return rv, io.EOF
	}
	if err != nil {
		return rv, err
	}
	res := NewStreamingPayloadResultWithViewsMethodUsertypeOK(&body)
	vres := &streamingpayloadresultwithviewsserviceviews.Usertype{res, s.view}
	if err := streamingpayloadresultwithviewsserviceviews.ValidateUsertype(vres); err != nil {
		return rv, goahttp.ErrValidationError("StreamingPayloadResultWithViewsService", "StreamingPayloadResultWithViewsMethod", err)
	}
	return streamingpayloadresultwithviewsservice.NewUsertype(vres), nil
}
`

var StreamingPayloadResultWithViewsClientStreamSetViewCode = `// SetView sets the view to render the float32 type before sending to the
// "StreamingPayloadResultWithViewsMethod" endpoint websocket connection.
func (s *StreamingPayloadResultWithViewsMethodClientStream) SetView(view string) {
	s.view = view
}
`

var StreamingPayloadResultWithExplicitViewServerStreamSendCode = `// SendAndClose streams instances of
// "streamingpayloadresultwithexplicitviewservice.Usertype" to the
// "StreamingPayloadResultWithExplicitViewMethod" endpoint websocket connection
// and closes the connection.
func (s *StreamingPayloadResultWithExplicitViewMethodServerStream) SendAndClose(v *streamingpayloadresultwithexplicitviewservice.Usertype) error {
	defer s.conn.Close()
	res := streamingpayloadresultwithexplicitviewservice.NewViewedUsertype(v, "extended")
	body := NewStreamingPayloadResultWithExplicitViewMethodResponseBodyExtended(res.Projected)
	return s.conn.WriteJSON(body)
}
`

var StreamingPayloadResultWithExplicitViewServerStreamRecvCode = `// Recv reads instances of "float32" from the
// "StreamingPayloadResultWithExplicitViewMethod" endpoint websocket connection.
func (s *StreamingPayloadResultWithExplicitViewMethodServerStream) Recv() (float32, error) {
	var (
		rv  float32
		msg *float32
		err error
	)
	// Upgrade the HTTP connection to a websocket connection only once. Connection
	// upgrade is done here so that authorization logic in the endpoint is executed
	// before calling the actual service method which may call Recv().
	s.once.Do(func() {
		var conn *websocket.Conn
		conn, err = s.upgrader.Upgrade(s.w, s.r, nil)
		if err != nil {
			return
		}
		if s.configurer != nil {
			conn = s.configurer(conn, s.cancel)
		}
		s.conn = conn
	})
	if err != nil {
		return rv, err
	}
	if err = s.conn.ReadJSON(&msg); err != nil {
		return rv, err
	}
	if msg == nil {
		return rv, io.EOF
	}
	return *msg, nil
}
`

var StreamingPayloadResultWithExplicitViewClientStreamSendCode = `// Send streams instances of "float32" to the
// "StreamingPayloadResultWithExplicitViewMethod" endpoint websocket connection.
func (s *StreamingPayloadResultWithExplicitViewMethodClientStream) Send(v float32) error {
	return s.conn.WriteJSON(v)
}
`

var StreamingPayloadResultWithExplicitViewClientStreamRecvCode = `// CloseAndRecv stops sending messages to the
// "StreamingPayloadResultWithExplicitViewMethod" endpoint websocket connection
// and reads instances of
// "streamingpayloadresultwithexplicitviewservice.Usertype" from the connection.
func (s *StreamingPayloadResultWithExplicitViewMethodClientStream) CloseAndRecv() (*streamingpayloadresultwithexplicitviewservice.Usertype, error) {
	var (
		rv   *streamingpayloadresultwithexplicitviewservice.Usertype
		body StreamingPayloadResultWithExplicitViewMethodResponseBody
		err  error
	)
	defer s.conn.Close()
	// Send a nil payload to the server implying end of message
	if err = s.conn.WriteJSON(nil); err != nil {
		return rv, err
	}
	err = s.conn.ReadJSON(&body)
	if websocket.IsCloseError(err, websocket.CloseNormalClosure) {
		s.conn.Close()
		return rv, io.EOF
	}
	if err != nil {
		return rv, err
	}
	res := NewStreamingPayloadResultWithExplicitViewMethodUsertypeOK(&body)
	vres := &streamingpayloadresultwithexplicitviewserviceviews.Usertype{res, "extended"}
	if err := streamingpayloadresultwithexplicitviewserviceviews.ValidateUsertype(vres); err != nil {
		return rv, goahttp.ErrValidationError("StreamingPayloadResultWithExplicitViewService", "StreamingPayloadResultWithExplicitViewMethod", err)
	}
	return streamingpayloadresultwithexplicitviewservice.NewUsertype(vres), nil
}
`

var StreamingPayloadResultListWithViewsServerStreamSendCode = `// SendAndClose streams instances of
// "streamingpayloadresultlistwithviewsservice.UsertypeList" to the
// "StreamingPayloadResultListWithViewsMethod" endpoint websocket
// connection and closes the connection.
func (s *StreamingPayloadResultListWithViewsMethodServerStream) SendAndClose(v streamingpayloadresultlistwithviewsservice.UsertypeList) error {
	defer s.conn.Close()
	res := streamingpayloadresultlistwithviewsservice.NewViewedUsertypeList(v, s.view)
	var body interface{}
	switch s.view {
	case "tiny":
		body = NewUsertypeResponseTinyList(res.Projected)
	case "extended":
		body = NewUsertypeResponseExtendedList(res.Projected)
	case "default", "":
		body = NewUsertypeResponseList(res.Projected)
	}
	return s.conn.WriteJSON(body)
}
`

var StreamingPayloadResultListWithViewsServerStreamRecvCode = `// Recv reads instances of "interface{}" from the
// "StreamingPayloadResultListWithViewsMethod" endpoint websocket
// connection.
func (s *StreamingPayloadResultListWithViewsMethodServerStream) Recv() (interface{}, error) {
	var (
		rv  interface{}
		msg *interface{}
		err error
	)
	// Upgrade the HTTP connection to a websocket connection only once. Connection
	// upgrade is done here so that authorization logic in the endpoint is executed
	// before calling the actual service method which may call Recv().
	s.once.Do(func() {
		var conn *websocket.Conn
		conn, err = s.upgrader.Upgrade(s.w, s.r, nil)
		if err != nil {
			return
		}
		if s.configurer != nil {
			conn = s.configurer(conn, s.cancel)
		}
		s.conn = conn
	})
	if err != nil {
		return rv, err
	}
	if err = s.conn.ReadJSON(&msg); err != nil {
		return rv, err
	}
	if msg == nil {
		return rv, io.EOF
	}
	return *msg, nil
}
`

var StreamingPayloadResultListWithViewsServerStreamSetViewCode = `// SetView sets the view to render the
// streamingpayloadresultlistwithviewsservice.UsertypeList type
// before sending to the "StreamingPayloadResultListWithViewsMethod"
// endpoint websocket connection.
func (s *StreamingPayloadResultListWithViewsMethodServerStream) SetView(view string) {
	s.view = view
}
`

var StreamingPayloadResultListWithViewsClientStreamSendCode = `// Send streams instances of "interface{}" to the
// "StreamingPayloadResultListWithViewsMethod" endpoint websocket
// connection.
func (s *StreamingPayloadResultListWithViewsMethodClientStream) Send(v interface{}) error {
	return s.conn.WriteJSON(v)
}
`

var StreamingPayloadResultListWithViewsClientStreamRecvCode = `// CloseAndRecv stops sending messages to the
// "StreamingPayloadResultListWithViewsMethod" endpoint websocket
// connection and reads instances of
// "streamingpayloadresultlistwithviewsservice.UsertypeList" from
// the connection.
func (s *StreamingPayloadResultListWithViewsMethodClientStream) CloseAndRecv() (streamingpayloadresultlistwithviewsservice.UsertypeList, error) {
	var (
		rv   streamingpayloadresultlistwithviewsservice.UsertypeList
		body StreamingPayloadResultListWithViewsMethodResponseBody
		err  error
	)
	defer s.conn.Close()
	// Send a nil payload to the server implying end of message
	if err = s.conn.WriteJSON(nil); err != nil {
		return rv, err
	}
	err = s.conn.ReadJSON(&body)
	if websocket.IsCloseError(err, websocket.CloseNormalClosure) {
		s.conn.Close()
		return rv, io.EOF
	}
	if err != nil {
		return rv, err
	}
	res := NewStreamingPayloadResultListWithViewsMethodUsertypeListOK(body)
	vres := streamingpayloadresultlistwithviewsserviceviews.UsertypeList{res, s.view}
	if err := streamingpayloadresultlistwithviewsserviceviews.ValidateUsertypeList(vres); err != nil {
		return rv, goahttp.ErrValidationError("StreamingPayloadResultListWithViewsService", "StreamingPayloadResultListWithViewsMethod", err)
	}
	return streamingpayloadresultlistwithviewsservice.NewUsertypeList(vres), nil
}
`

var StreamingPayloadResultListWithViewsClientStreamSetViewCode = `// SetView sets the view to render the interface{} type before sending to the
// "StreamingPayloadResultListWithViewsMethod" endpoint websocket
// connection.
func (s *StreamingPayloadResultListWithViewsMethodClientStream) SetView(view string) {
	s.view = view
}
`

var StreamingPayloadResultListWithExplicitViewServerStreamSendCode = `// SendAndClose streams instances of
// "streamingpayloadresultlistwithexplicitviewservice.UsertypeList"
// to the "StreamingPayloadResultListWithExplicitViewMethod" endpoint
// websocket connection and closes the connection.
func (s *StreamingPayloadResultListWithExplicitViewMethodServerStream) SendAndClose(v streamingpayloadresultlistwithexplicitviewservice.UsertypeList) error {
	defer s.conn.Close()
	res := streamingpayloadresultlistwithexplicitviewservice.NewViewedUsertypeList(v, "tiny")
	body := NewUsertypeResponseTinyList(res.Projected)
	return s.conn.WriteJSON(body)
}
`

var StreamingPayloadResultListWithExplicitViewServerStreamRecvCode = `// Recv reads instances of "interface{}" from the
// "StreamingPayloadResultListWithExplicitViewMethod" endpoint websocket
// connection.
func (s *StreamingPayloadResultListWithExplicitViewMethodServerStream) Recv() (interface{}, error) {
	var (
		rv  interface{}
		msg *interface{}
		err error
	)
	// Upgrade the HTTP connection to a websocket connection only once. Connection
	// upgrade is done here so that authorization logic in the endpoint is executed
	// before calling the actual service method which may call Recv().
	s.once.Do(func() {
		var conn *websocket.Conn
		conn, err = s.upgrader.Upgrade(s.w, s.r, nil)
		if err != nil {
			return
		}
		if s.configurer != nil {
			conn = s.configurer(conn, s.cancel)
		}
		s.conn = conn
	})
	if err != nil {
		return rv, err
	}
	if err = s.conn.ReadJSON(&msg); err != nil {
		return rv, err
	}
	if msg == nil {
		return rv, io.EOF
	}
	return *msg, nil
}
`

var StreamingPayloadResultListWithExplicitViewClientStreamSendCode = `// Send streams instances of "interface{}" to the
// "StreamingPayloadResultListWithExplicitViewMethod" endpoint websocket
// connection.
func (s *StreamingPayloadResultListWithExplicitViewMethodClientStream) Send(v interface{}) error {
	return s.conn.WriteJSON(v)
}
`

var StreamingPayloadResultListWithExplicitViewClientStreamRecvCode = `// CloseAndRecv stops sending messages to the
// "StreamingPayloadResultListWithExplicitViewMethod" endpoint websocket
// connection and reads instances of
// "streamingpayloadresultlistwithexplicitviewservice.UsertypeList"
// from the connection.
func (s *StreamingPayloadResultListWithExplicitViewMethodClientStream) CloseAndRecv() (streamingpayloadresultlistwithexplicitviewservice.UsertypeList, error) {
	var (
		rv   streamingpayloadresultlistwithexplicitviewservice.UsertypeList
		body StreamingPayloadResultListWithExplicitViewMethodResponseBody
		err  error
	)
	defer s.conn.Close()
	// Send a nil payload to the server implying end of message
	if err = s.conn.WriteJSON(nil); err != nil {
		return rv, err
	}
	err = s.conn.ReadJSON(&body)
	if websocket.IsCloseError(err, websocket.CloseNormalClosure) {
		s.conn.Close()
		return rv, io.EOF
	}
	if err != nil {
		return rv, err
	}
	res := NewStreamingPayloadResultListWithExplicitViewMethodUsertypeListOK(body)
	vres := streamingpayloadresultlistwithexplicitviewserviceviews.UsertypeList{res, "tiny"}
	if err := streamingpayloadresultlistwithexplicitviewserviceviews.ValidateUsertypeList(vres); err != nil {
		return rv, goahttp.ErrValidationError("StreamingPayloadResultListWithExplicitViewService", "StreamingPayloadResultListWithExplicitViewMethod", err)
	}
	return streamingpayloadresultlistwithexplicitviewservice.NewUsertypeList(vres), nil
}
`

var StreamingPayloadPrimitiveServerStreamSendCode = `// SendAndClose streams instances of "string" to the
// "StreamingPayloadPrimitiveMethod" endpoint websocket connection and closes
// the connection.
func (s *StreamingPayloadPrimitiveMethodServerStream) SendAndClose(v string) error {
	defer s.conn.Close()
	res := v
	return s.conn.WriteJSON(res)
}
`

var StreamingPayloadPrimitiveServerStreamRecvCode = `// Recv reads instances of "string" from the "StreamingPayloadPrimitiveMethod"
// endpoint websocket connection.
func (s *StreamingPayloadPrimitiveMethodServerStream) Recv() (string, error) {
	var (
		rv  string
		msg *string
		err error
	)
	// Upgrade the HTTP connection to a websocket connection only once. Connection
	// upgrade is done here so that authorization logic in the endpoint is executed
	// before calling the actual service method which may call Recv().
	s.once.Do(func() {
		var conn *websocket.Conn
		conn, err = s.upgrader.Upgrade(s.w, s.r, nil)
		if err != nil {
			return
		}
		if s.configurer != nil {
			conn = s.configurer(conn, s.cancel)
		}
		s.conn = conn
	})
	if err != nil {
		return rv, err
	}
	if err = s.conn.ReadJSON(&msg); err != nil {
		return rv, err
	}
	if msg == nil {
		return rv, io.EOF
	}
	return *msg, nil
}
`

var StreamingPayloadPrimitiveClientStreamSendCode = `// Send streams instances of "string" to the "StreamingPayloadPrimitiveMethod"
// endpoint websocket connection.
func (s *StreamingPayloadPrimitiveMethodClientStream) Send(v string) error {
	return s.conn.WriteJSON(v)
}
`

var StreamingPayloadPrimitiveClientStreamRecvCode = `// CloseAndRecv stops sending messages to the "StreamingPayloadPrimitiveMethod"
// endpoint websocket connection and reads instances of "string" from the
// connection.
func (s *StreamingPayloadPrimitiveMethodClientStream) CloseAndRecv() (string, error) {
	var (
		rv   string
		body string
		err  error
	)
	defer s.conn.Close()
	// Send a nil payload to the server implying end of message
	if err = s.conn.WriteJSON(nil); err != nil {
		return rv, err
	}
	err = s.conn.ReadJSON(&body)
	if websocket.IsCloseError(err, websocket.CloseNormalClosure) {
		s.conn.Close()
		return rv, io.EOF
	}
	if err != nil {
		return rv, err
	}
	return body, nil
}
`

var StreamingPayloadPrimitiveArrayServerStreamSendCode = `// SendAndClose streams instances of "[]string" to the
// "StreamingPayloadPrimitiveArrayMethod" endpoint websocket connection and
// closes the connection.
func (s *StreamingPayloadPrimitiveArrayMethodServerStream) SendAndClose(v []string) error {
	defer s.conn.Close()
	res := v
	return s.conn.WriteJSON(res)
}
`

var StreamingPayloadPrimitiveArrayServerStreamRecvCode = `// Recv reads instances of "[]int32" from the
// "StreamingPayloadPrimitiveArrayMethod" endpoint websocket connection.
func (s *StreamingPayloadPrimitiveArrayMethodServerStream) Recv() ([]int32, error) {
	var (
		rv   []int32
		body []int32
		err  error
	)
	// Upgrade the HTTP connection to a websocket connection only once. Connection
	// upgrade is done here so that authorization logic in the endpoint is executed
	// before calling the actual service method which may call Recv().
	s.once.Do(func() {
		var conn *websocket.Conn
		conn, err = s.upgrader.Upgrade(s.w, s.r, nil)
		if err != nil {
			return
		}
		if s.configurer != nil {
			conn = s.configurer(conn, s.cancel)
		}
		s.conn = conn
	})
	if err != nil {
		return rv, err
	}
	if err = s.conn.ReadJSON(&body); err != nil {
		return rv, err
	}
	if body == nil {
		return rv, io.EOF
	}
	return body, nil
}
`

var StreamingPayloadPrimitiveArrayClientStreamSendCode = `// Send streams instances of "[]int32" to the
// "StreamingPayloadPrimitiveArrayMethod" endpoint websocket connection.
func (s *StreamingPayloadPrimitiveArrayMethodClientStream) Send(v []int32) error {
	return s.conn.WriteJSON(v)
}
`

var StreamingPayloadPrimitiveArrayClientStreamRecvCode = `// CloseAndRecv stops sending messages to the
// "StreamingPayloadPrimitiveArrayMethod" endpoint websocket connection and
// reads instances of "[]string" from the connection.
func (s *StreamingPayloadPrimitiveArrayMethodClientStream) CloseAndRecv() ([]string, error) {
	var (
		rv   []string
		body []string
		err  error
	)
	defer s.conn.Close()
	// Send a nil payload to the server implying end of message
	if err = s.conn.WriteJSON(nil); err != nil {
		return rv, err
	}
	err = s.conn.ReadJSON(&body)
	if websocket.IsCloseError(err, websocket.CloseNormalClosure) {
		s.conn.Close()
		return rv, io.EOF
	}
	if err != nil {
		return rv, err
	}
	return body, nil
}
`

var StreamingPayloadPrimitiveMapServerStreamSendCode = `// SendAndClose streams instances of "map[int]int" to the
// "StreamingPayloadPrimitiveMapMethod" endpoint websocket connection and
// closes the connection.
func (s *StreamingPayloadPrimitiveMapMethodServerStream) SendAndClose(v map[int]int) error {
	defer s.conn.Close()
	res := v
	return s.conn.WriteJSON(res)
}
`

var StreamingPayloadPrimitiveMapServerStreamRecvCode = `// Recv reads instances of "map[string]int32" from the
// "StreamingPayloadPrimitiveMapMethod" endpoint websocket connection.
func (s *StreamingPayloadPrimitiveMapMethodServerStream) Recv() (map[string]int32, error) {
	var (
		rv   map[string]int32
		body map[string]int32
		err  error
	)
	// Upgrade the HTTP connection to a websocket connection only once. Connection
	// upgrade is done here so that authorization logic in the endpoint is executed
	// before calling the actual service method which may call Recv().
	s.once.Do(func() {
		var conn *websocket.Conn
		conn, err = s.upgrader.Upgrade(s.w, s.r, nil)
		if err != nil {
			return
		}
		if s.configurer != nil {
			conn = s.configurer(conn, s.cancel)
		}
		s.conn = conn
	})
	if err != nil {
		return rv, err
	}
	if err = s.conn.ReadJSON(&body); err != nil {
		return rv, err
	}
	if body == nil {
		return rv, io.EOF
	}
	return body, nil
}
`

var StreamingPayloadPrimitiveMapClientStreamSendCode = `// Send streams instances of "map[string]int32" to the
// "StreamingPayloadPrimitiveMapMethod" endpoint websocket connection.
func (s *StreamingPayloadPrimitiveMapMethodClientStream) Send(v map[string]int32) error {
	return s.conn.WriteJSON(v)
}
`

var StreamingPayloadPrimitiveMapClientStreamRecvCode = `// CloseAndRecv stops sending messages to the
// "StreamingPayloadPrimitiveMapMethod" endpoint websocket connection and reads
// instances of "map[int]int" from the connection.
func (s *StreamingPayloadPrimitiveMapMethodClientStream) CloseAndRecv() (map[int]int, error) {
	var (
		rv   map[int]int
		body map[int]int
		err  error
	)
	defer s.conn.Close()
	// Send a nil payload to the server implying end of message
	if err = s.conn.WriteJSON(nil); err != nil {
		return rv, err
	}
	err = s.conn.ReadJSON(&body)
	if websocket.IsCloseError(err, websocket.CloseNormalClosure) {
		s.conn.Close()
		return rv, io.EOF
	}
	if err != nil {
		return rv, err
	}
	return body, nil
}
`

var StreamingPayloadUserTypeArrayServerStreamSendCode = `// SendAndClose streams instances of "string" to the
// "StreamingPayloadUserTypeArrayMethod" endpoint websocket connection and
// closes the connection.
func (s *StreamingPayloadUserTypeArrayMethodServerStream) SendAndClose(v string) error {
	defer s.conn.Close()
	res := v
	return s.conn.WriteJSON(res)
}
`

var StreamingPayloadUserTypeArrayServerStreamRecvCode = `// Recv reads instances of
// "[]*streamingpayloadusertypearrayservice.RequestType" from the
// "StreamingPayloadUserTypeArrayMethod" endpoint websocket connection.
func (s *StreamingPayloadUserTypeArrayMethodServerStream) Recv() ([]*streamingpayloadusertypearrayservice.RequestType, error) {
	var (
		rv   []*streamingpayloadusertypearrayservice.RequestType
		body []*RequestType
		err  error
	)
	// Upgrade the HTTP connection to a websocket connection only once. Connection
	// upgrade is done here so that authorization logic in the endpoint is executed
	// before calling the actual service method which may call Recv().
	s.once.Do(func() {
		var conn *websocket.Conn
		conn, err = s.upgrader.Upgrade(s.w, s.r, nil)
		if err != nil {
			return
		}
		if s.configurer != nil {
			conn = s.configurer(conn, s.cancel)
		}
		s.conn = conn
	})
	if err != nil {
		return rv, err
	}
	if err = s.conn.ReadJSON(&body); err != nil {
		return rv, err
	}
	if body == nil {
		return rv, io.EOF
	}
	return NewStreamingPayloadUserTypeArrayMethodArray(body), nil
}
`

var StreamingPayloadUserTypeArrayClientStreamSendCode = `// Send streams instances of
// "[]*streamingpayloadusertypearrayservice.RequestType" to the
// "StreamingPayloadUserTypeArrayMethod" endpoint websocket connection.
func (s *StreamingPayloadUserTypeArrayMethodClientStream) Send(v []*streamingpayloadusertypearrayservice.RequestType) error {
	body := NewRequestType(v)
	return s.conn.WriteJSON(body)
}
`

var StreamingPayloadUserTypeArrayClientStreamRecvCode = `// CloseAndRecv stops sending messages to the
// "StreamingPayloadUserTypeArrayMethod" endpoint websocket connection and
// reads instances of "string" from the connection.
func (s *StreamingPayloadUserTypeArrayMethodClientStream) CloseAndRecv() (string, error) {
	var (
		rv   string
		body string
		err  error
	)
	defer s.conn.Close()
	// Send a nil payload to the server implying end of message
	if err = s.conn.WriteJSON(nil); err != nil {
		return rv, err
	}
	err = s.conn.ReadJSON(&body)
	if websocket.IsCloseError(err, websocket.CloseNormalClosure) {
		s.conn.Close()
		return rv, io.EOF
	}
	if err != nil {
		return rv, err
	}
	return body, nil
}
`

var StreamingPayloadUserTypeMapServerStreamSendCode = `// SendAndClose streams instances of "[]string" to the
// "StreamingPayloadUserTypeMapMethod" endpoint websocket connection and closes
// the connection.
func (s *StreamingPayloadUserTypeMapMethodServerStream) SendAndClose(v []string) error {
	defer s.conn.Close()
	res := v
	return s.conn.WriteJSON(res)
}
`

var StreamingPayloadUserTypeMapServerStreamRecvCode = `// Recv reads instances of
// "map[string]*streamingpayloadusertypemapservice.RequestType" from the
// "StreamingPayloadUserTypeMapMethod" endpoint websocket connection.
func (s *StreamingPayloadUserTypeMapMethodServerStream) Recv() (map[string]*streamingpayloadusertypemapservice.RequestType, error) {
	var (
		rv   map[string]*streamingpayloadusertypemapservice.RequestType
		body map[string]*RequestType
		err  error
	)
	// Upgrade the HTTP connection to a websocket connection only once. Connection
	// upgrade is done here so that authorization logic in the endpoint is executed
	// before calling the actual service method which may call Recv().
	s.once.Do(func() {
		var conn *websocket.Conn
		conn, err = s.upgrader.Upgrade(s.w, s.r, nil)
		if err != nil {
			return
		}
		if s.configurer != nil {
			conn = s.configurer(conn, s.cancel)
		}
		s.conn = conn
	})
	if err != nil {
		return rv, err
	}
	if err = s.conn.ReadJSON(&body); err != nil {
		return rv, err
	}
	if body == nil {
		return rv, io.EOF
	}
	return NewStreamingPayloadUserTypeMapMethodMap(body), nil
}
`

var StreamingPayloadUserTypeMapClientStreamSendCode = `// Send streams instances of
// "map[string]*streamingpayloadusertypemapservice.RequestType" to the
// "StreamingPayloadUserTypeMapMethod" endpoint websocket connection.
func (s *StreamingPayloadUserTypeMapMethodClientStream) Send(v map[string]*streamingpayloadusertypemapservice.RequestType) error {
	body := NewMapStringRequestType(v)
	return s.conn.WriteJSON(body)
}
`

var StreamingPayloadUserTypeMapClientStreamRecvCode = `// CloseAndRecv stops sending messages to the
// "StreamingPayloadUserTypeMapMethod" endpoint websocket connection and reads
// instances of "[]string" from the connection.
func (s *StreamingPayloadUserTypeMapMethodClientStream) CloseAndRecv() ([]string, error) {
	var (
		rv   []string
		body []string
		err  error
	)
	defer s.conn.Close()
	// Send a nil payload to the server implying end of message
	if err = s.conn.WriteJSON(nil); err != nil {
		return rv, err
	}
	err = s.conn.ReadJSON(&body)
	if websocket.IsCloseError(err, websocket.CloseNormalClosure) {
		s.conn.Close()
		return rv, io.EOF
	}
	if err != nil {
		return rv, err
	}
	return body, nil
}
`

var BidirectionalStreamingServerHandlerInitCode = `// NewBidirectionalStreamingMethodHandler creates a HTTP handler which loads
// the HTTP request and calls the "BidirectionalStreamingService" service
// "BidirectionalStreamingMethod" endpoint.
func NewBidirectionalStreamingMethodHandler(
	endpoint goa.Endpoint,
	mux goahttp.Muxer,
	decoder func(*http.Request) goahttp.Decoder,
	encoder func(context.Context, http.ResponseWriter) goahttp.Encoder,
	errhandler func(context.Context, http.ResponseWriter, error),
	formatter func(err error) goahttp.Statuser,
	upgrader goahttp.Upgrader,
	configurer goahttp.ConnConfigureFunc,
) http.Handler {
	var (
		decodeRequest = DecodeBidirectionalStreamingMethodRequest(mux, decoder)
		encodeError   = goahttp.ErrorEncoder(encoder, formatter)
	)
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := context.WithValue(r.Context(), goahttp.AcceptTypeKey, r.Header.Get("Accept"))
		ctx = context.WithValue(ctx, goa.MethodKey, "BidirectionalStreamingMethod")
		ctx = context.WithValue(ctx, goa.ServiceKey, "BidirectionalStreamingService")
		payload, err := decodeRequest(r)
		if err != nil {
			if err := encodeError(ctx, w, err); err != nil {
				errhandler(ctx, w, err)
			}
			return
		}
		var cancel context.CancelFunc
		ctx, cancel = context.WithCancel(ctx)
		v := &bidirectionalstreamingservice.BidirectionalStreamingMethodEndpointInput{
			Stream: &BidirectionalStreamingMethodServerStream{
				upgrader:   upgrader,
				configurer: configurer,
				cancel:     cancel,
				w:          w,
				r:          r,
			},
			Payload: payload.(*bidirectionalstreamingservice.Payload),
		}
		_, err = endpoint(ctx, v)
		if err != nil {
			if _, werr := w.Write(nil); werr == http.ErrHijacked {
				// Response writer has been hijacked, do not encode the error
				errhandler(ctx, w, err)
				return
			}
			if err := encodeError(ctx, w, err); err != nil {
				errhandler(ctx, w, err)
			}
			return
		}
	})
}
`

var BidirectionalStreamingServerStreamSendCode = `// Send streams instances of "bidirectionalstreamingservice.UserType" to the
// "BidirectionalStreamingMethod" endpoint websocket connection.
func (s *BidirectionalStreamingMethodServerStream) Send(v *bidirectionalstreamingservice.UserType) error {
	var err error
	// Upgrade the HTTP connection to a websocket connection only once. Connection
	// upgrade is done here so that authorization logic in the endpoint is executed
	// before calling the actual service method which may call Send().
	s.once.Do(func() {
		var conn *websocket.Conn
		conn, err = s.upgrader.Upgrade(s.w, s.r, nil)
		if err != nil {
			return
		}
		if s.configurer != nil {
			conn = s.configurer(conn, s.cancel)
		}
		s.conn = conn
	})
	if err != nil {
		return err
	}
	res := v
	body := NewBidirectionalStreamingMethodResponseBody(res)
	return s.conn.WriteJSON(body)
}
`

var BidirectionalStreamingServerStreamRecvCode = `// Recv reads instances of "bidirectionalstreamingservice.Request" from the
// "BidirectionalStreamingMethod" endpoint websocket connection.
func (s *BidirectionalStreamingMethodServerStream) Recv() (*bidirectionalstreamingservice.Request, error) {
	var (
		rv  *bidirectionalstreamingservice.Request
		msg *BidirectionalStreamingMethodStreamingBody
		err error
	)
	// Upgrade the HTTP connection to a websocket connection only once. Connection
	// upgrade is done here so that authorization logic in the endpoint is executed
	// before calling the actual service method which may call Recv().
	s.once.Do(func() {
		var conn *websocket.Conn
		conn, err = s.upgrader.Upgrade(s.w, s.r, nil)
		if err != nil {
			return
		}
		if s.configurer != nil {
			conn = s.configurer(conn, s.cancel)
		}
		s.conn = conn
	})
	if err != nil {
		return rv, err
	}
	if err = s.conn.ReadJSON(&msg); err != nil {
		return rv, err
	}
	if msg == nil {
		return rv, io.EOF
	}
	return NewBidirectionalStreamingMethodStreamingBody(msg), nil
}
`

var BidirectionalStreamingServerStreamCloseCode = `// Close closes the "BidirectionalStreamingMethod" endpoint websocket
// connection.
func (s *BidirectionalStreamingMethodServerStream) Close() error {
	var err error
	if s.conn == nil {
		return nil
	}
	if err = s.conn.WriteControl(
		websocket.CloseMessage,
		websocket.FormatCloseMessage(websocket.CloseNormalClosure, "server closing connection"),
		time.Now().Add(time.Second),
	); err != nil {
		return err
	}
	return s.conn.Close()
}
`

var BidirectionalStreamingClientEndpointCode = `// BidirectionalStreamingMethod returns an endpoint that makes HTTP requests to
// the BidirectionalStreamingService service BidirectionalStreamingMethod
// server.
func (c *Client) BidirectionalStreamingMethod() goa.Endpoint {
	var (
		encodeRequest  = EncodeBidirectionalStreamingMethodRequest(c.encoder)
		decodeResponse = DecodeBidirectionalStreamingMethodResponse(c.decoder, c.RestoreResponseBody)
	)
	return func(ctx context.Context, v interface{}) (interface{}, error) {
		req, err := c.BuildBidirectionalStreamingMethodRequest(ctx, v)
		if err != nil {
			return nil, err
		}
		err = encodeRequest(req, v)
		if err != nil {
			return nil, err
		}
		var cancel context.CancelFunc
		ctx, cancel = context.WithCancel(ctx)
		conn, resp, err := c.dialer.DialContext(ctx, req.URL.String(), req.Header)
		if err != nil {
			if resp != nil {
				return decodeResponse(resp)
			}
			return nil, goahttp.ErrRequestError("BidirectionalStreamingService", "BidirectionalStreamingMethod", err)
		}
		if c.configurer.BidirectionalStreamingMethodFn != nil {
			conn = c.configurer.BidirectionalStreamingMethodFn(conn, cancel)
		}
		stream := &BidirectionalStreamingMethodClientStream{conn: conn}
		return stream, nil
	}
}
`

var BidirectionalStreamingClientStreamSendCode = `// Send streams instances of "bidirectionalstreamingservice.Request" to the
// "BidirectionalStreamingMethod" endpoint websocket connection.
func (s *BidirectionalStreamingMethodClientStream) Send(v *bidirectionalstreamingservice.Request) error {
	body := NewBidirectionalStreamingMethodStreamingBody(v)
	return s.conn.WriteJSON(body)
}
`

var BidirectionalStreamingClientStreamRecvCode = `// Recv reads instances of "bidirectionalstreamingservice.UserType" from the
// "BidirectionalStreamingMethod" endpoint websocket connection.
func (s *BidirectionalStreamingMethodClientStream) Recv() (*bidirectionalstreamingservice.UserType, error) {
	var (
		rv   *bidirectionalstreamingservice.UserType
		body BidirectionalStreamingMethodResponseBody
		err  error
	)
	err = s.conn.ReadJSON(&body)
	if websocket.IsCloseError(err, websocket.CloseNormalClosure) {
		return rv, io.EOF
	}
	if err != nil {
		return rv, err
	}
	res := NewBidirectionalStreamingMethodUserTypeOK(&body)
	return res, nil
}
`

var BidirectionalStreamingClientStreamCloseCode = `// Close closes the "BidirectionalStreamingMethod" endpoint websocket
// connection.
func (s *BidirectionalStreamingMethodClientStream) Close() error {
	var err error
	// Send a nil payload to the server implying client closing connection.
	if err = s.conn.WriteJSON(nil); err != nil {
		return err
	}
	return s.conn.Close()
}
`

var BidirectionalStreamingNoPayloadServerHandlerInitCode = `// NewBidirectionalStreamingNoPayloadMethodHandler creates a HTTP handler which
// loads the HTTP request and calls the
// "BidirectionalStreamingNoPayloadService" service
// "BidirectionalStreamingNoPayloadMethod" endpoint.
func NewBidirectionalStreamingNoPayloadMethodHandler(
	endpoint goa.Endpoint,
	mux goahttp.Muxer,
	decoder func(*http.Request) goahttp.Decoder,
	encoder func(context.Context, http.ResponseWriter) goahttp.Encoder,
	errhandler func(context.Context, http.ResponseWriter, error),
	formatter func(err error) goahttp.Statuser,
	upgrader goahttp.Upgrader,
	configurer goahttp.ConnConfigureFunc,
) http.Handler {
	var (
		encodeError = goahttp.ErrorEncoder(encoder, formatter)
	)
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := context.WithValue(r.Context(), goahttp.AcceptTypeKey, r.Header.Get("Accept"))
		ctx = context.WithValue(ctx, goa.MethodKey, "BidirectionalStreamingNoPayloadMethod")
		ctx = context.WithValue(ctx, goa.ServiceKey, "BidirectionalStreamingNoPayloadService")
		var err error
		var cancel context.CancelFunc
		ctx, cancel = context.WithCancel(ctx)
		v := &bidirectionalstreamingnopayloadservice.BidirectionalStreamingNoPayloadMethodEndpointInput{
			Stream: &BidirectionalStreamingNoPayloadMethodServerStream{
				upgrader:   upgrader,
				configurer: configurer,
				cancel:     cancel,
				w:          w,
				r:          r,
			},
		}
		_, err = endpoint(ctx, v)
		if err != nil {
			if _, werr := w.Write(nil); werr == http.ErrHijacked {
				// Response writer has been hijacked, do not encode the error
				errhandler(ctx, w, err)
				return
			}
			if err := encodeError(ctx, w, err); err != nil {
				errhandler(ctx, w, err)
			}
			return
		}
	})
}
`

var BidirectionalStreamingNoPayloadServerStreamCloseCode = `// Close closes the "BidirectionalStreamingNoPayloadMethod" endpoint websocket
// connection.
func (s *BidirectionalStreamingNoPayloadMethodServerStream) Close() error {
	var err error
	if s.conn == nil {
		return nil
	}
	if err = s.conn.WriteControl(
		websocket.CloseMessage,
		websocket.FormatCloseMessage(websocket.CloseNormalClosure, "server closing connection"),
		time.Now().Add(time.Second),
	); err != nil {
		return err
	}
	return s.conn.Close()
}
`

var BidirectionalStreamingNoPayloadClientEndpointCode = `// BidirectionalStreamingNoPayloadMethod returns an endpoint that makes HTTP
// requests to the BidirectionalStreamingNoPayloadService service
// BidirectionalStreamingNoPayloadMethod server.
func (c *Client) BidirectionalStreamingNoPayloadMethod() goa.Endpoint {
	var (
		decodeResponse = DecodeBidirectionalStreamingNoPayloadMethodResponse(c.decoder, c.RestoreResponseBody)
	)
	return func(ctx context.Context, v interface{}) (interface{}, error) {
		req, err := c.BuildBidirectionalStreamingNoPayloadMethodRequest(ctx, v)
		if err != nil {
			return nil, err
		}
		var cancel context.CancelFunc
		ctx, cancel = context.WithCancel(ctx)
		conn, resp, err := c.dialer.DialContext(ctx, req.URL.String(), req.Header)
		if err != nil {
			if resp != nil {
				return decodeResponse(resp)
			}
			return nil, goahttp.ErrRequestError("BidirectionalStreamingNoPayloadService", "BidirectionalStreamingNoPayloadMethod", err)
		}
		if c.configurer.BidirectionalStreamingNoPayloadMethodFn != nil {
			conn = c.configurer.BidirectionalStreamingNoPayloadMethodFn(conn, cancel)
		}
		stream := &BidirectionalStreamingNoPayloadMethodClientStream{conn: conn}
		return stream, nil
	}
}
`

var BidirectionalStreamingNoPayloadClientStreamSendCode = `// Send streams instances of "bidirectionalstreamingnopayloadservice.Request"
// to the "BidirectionalStreamingNoPayloadMethod" endpoint websocket connection.
func (s *BidirectionalStreamingNoPayloadMethodClientStream) Send(v *bidirectionalstreamingnopayloadservice.Request) error {
	body := NewBidirectionalStreamingNoPayloadMethodStreamingBody(v)
	return s.conn.WriteJSON(body)
}
`

var BidirectionalStreamingNoPayloadClientStreamRecvCode = `// Recv reads instances of "bidirectionalstreamingnopayloadservice.UserType"
// from the "BidirectionalStreamingNoPayloadMethod" endpoint websocket
// connection.
func (s *BidirectionalStreamingNoPayloadMethodClientStream) Recv() (*bidirectionalstreamingnopayloadservice.UserType, error) {
	var (
		rv   *bidirectionalstreamingnopayloadservice.UserType
		body BidirectionalStreamingNoPayloadMethodResponseBody
		err  error
	)
	err = s.conn.ReadJSON(&body)
	if websocket.IsCloseError(err, websocket.CloseNormalClosure) {
		return rv, io.EOF
	}
	if err != nil {
		return rv, err
	}
	res := NewBidirectionalStreamingNoPayloadMethodUserTypeOK(&body)
	return res, nil
}
`

var BidirectionalStreamingNoPayloadClientStreamCloseCode = `// Close closes the "BidirectionalStreamingNoPayloadMethod" endpoint websocket
// connection.
func (s *BidirectionalStreamingNoPayloadMethodClientStream) Close() error {
	var err error
	// Send a nil payload to the server implying client closing connection.
	if err = s.conn.WriteJSON(nil); err != nil {
		return err
	}
	return s.conn.Close()
}
`

var BidirectionalStreamingResultWithViewsServerStreamSendCode = `// Send streams instances of
// "bidirectionalstreamingresultwithviewsservice.Usertype" to the
// "BidirectionalStreamingResultWithViewsMethod" endpoint websocket connection.
func (s *BidirectionalStreamingResultWithViewsMethodServerStream) Send(v *bidirectionalstreamingresultwithviewsservice.Usertype) error {
	var err error
	// Upgrade the HTTP connection to a websocket connection only once. Connection
	// upgrade is done here so that authorization logic in the endpoint is executed
	// before calling the actual service method which may call Send().
	s.once.Do(func() {
		respHdr := make(http.Header)
		respHdr.Add("goa-view", s.view)
		var conn *websocket.Conn
		conn, err = s.upgrader.Upgrade(s.w, s.r, respHdr)
		if err != nil {
			return
		}
		if s.configurer != nil {
			conn = s.configurer(conn, s.cancel)
		}
		s.conn = conn
	})
	if err != nil {
		return err
	}
	res := bidirectionalstreamingresultwithviewsservice.NewViewedUsertype(v, s.view)
	var body interface{}
	switch s.view {
	case "tiny":
		body = NewBidirectionalStreamingResultWithViewsMethodResponseBodyTiny(res.Projected)
	case "extended":
		body = NewBidirectionalStreamingResultWithViewsMethodResponseBodyExtended(res.Projected)
	case "default", "":
		body = NewBidirectionalStreamingResultWithViewsMethodResponseBody(res.Projected)
	}
	return s.conn.WriteJSON(body)
}
`

var BidirectionalStreamingResultWithViewsServerStreamRecvCode = `// Recv reads instances of "float32" from the
// "BidirectionalStreamingResultWithViewsMethod" endpoint websocket connection.
func (s *BidirectionalStreamingResultWithViewsMethodServerStream) Recv() (float32, error) {
	var (
		rv  float32
		msg *float32
		err error
	)
	// Upgrade the HTTP connection to a websocket connection only once. Connection
	// upgrade is done here so that authorization logic in the endpoint is executed
	// before calling the actual service method which may call Recv().
	s.once.Do(func() {
		var conn *websocket.Conn
		conn, err = s.upgrader.Upgrade(s.w, s.r, nil)
		if err != nil {
			return
		}
		if s.configurer != nil {
			conn = s.configurer(conn, s.cancel)
		}
		s.conn = conn
	})
	if err != nil {
		return rv, err
	}
	if err = s.conn.ReadJSON(&msg); err != nil {
		return rv, err
	}
	if msg == nil {
		return rv, io.EOF
	}
	return *msg, nil
}
`

var BidirectionalStreamingResultWithViewsServerStreamCloseCode = `// Close closes the "BidirectionalStreamingResultWithViewsMethod" endpoint
// websocket connection.
func (s *BidirectionalStreamingResultWithViewsMethodServerStream) Close() error {
	var err error
	if s.conn == nil {
		return nil
	}
	if err = s.conn.WriteControl(
		websocket.CloseMessage,
		websocket.FormatCloseMessage(websocket.CloseNormalClosure, "server closing connection"),
		time.Now().Add(time.Second),
	); err != nil {
		return err
	}
	return s.conn.Close()
}
`

var BidirectionalStreamingResultWithViewsServerStreamSetViewCode = `// SetView sets the view to render the
// bidirectionalstreamingresultwithviewsservice.Usertype type before sending to
// the "BidirectionalStreamingResultWithViewsMethod" endpoint websocket
// connection.
func (s *BidirectionalStreamingResultWithViewsMethodServerStream) SetView(view string) {
	s.view = view
}
`

var BidirectionalStreamingResultWithViewsClientStreamSendCode = `// Send streams instances of "float32" to the
// "BidirectionalStreamingResultWithViewsMethod" endpoint websocket connection.
func (s *BidirectionalStreamingResultWithViewsMethodClientStream) Send(v float32) error {
	return s.conn.WriteJSON(v)
}
`

var BidirectionalStreamingResultWithViewsClientStreamRecvCode = `// Recv reads instances of
// "bidirectionalstreamingresultwithviewsservice.Usertype" from the
// "BidirectionalStreamingResultWithViewsMethod" endpoint websocket connection.
func (s *BidirectionalStreamingResultWithViewsMethodClientStream) Recv() (*bidirectionalstreamingresultwithviewsservice.Usertype, error) {
	var (
		rv   *bidirectionalstreamingresultwithviewsservice.Usertype
		body BidirectionalStreamingResultWithViewsMethodResponseBody
		err  error
	)
	err = s.conn.ReadJSON(&body)
	if websocket.IsCloseError(err, websocket.CloseNormalClosure) {
		return rv, io.EOF
	}
	if err != nil {
		return rv, err
	}
	res := NewBidirectionalStreamingResultWithViewsMethodUsertypeOK(&body)
	vres := &bidirectionalstreamingresultwithviewsserviceviews.Usertype{res, s.view}
	if err := bidirectionalstreamingresultwithviewsserviceviews.ValidateUsertype(vres); err != nil {
		return rv, goahttp.ErrValidationError("BidirectionalStreamingResultWithViewsService", "BidirectionalStreamingResultWithViewsMethod", err)
	}
	return bidirectionalstreamingresultwithviewsservice.NewUsertype(vres), nil
}
`

var BidirectionalStreamingResultWithViewsClientStreamCloseCode = `// Close closes the "BidirectionalStreamingResultWithViewsMethod" endpoint
// websocket connection.
func (s *BidirectionalStreamingResultWithViewsMethodClientStream) Close() error {
	var err error
	// Send a nil payload to the server implying client closing connection.
	if err = s.conn.WriteJSON(nil); err != nil {
		return err
	}
	return s.conn.Close()
}
`

var BidirectionalStreamingResultWithViewsClientStreamSetViewCode = `// SetView sets the view to render the float32 type before sending to the
// "BidirectionalStreamingResultWithViewsMethod" endpoint websocket connection.
func (s *BidirectionalStreamingResultWithViewsMethodClientStream) SetView(view string) {
	s.view = view
}
`

var BidirectionalStreamingResultWithExplicitViewServerStreamSendCode = `// Send streams instances of
// "bidirectionalstreamingresultwithexplicitviewservice.Usertype" to the
// "BidirectionalStreamingResultWithExplicitViewMethod" endpoint websocket
// connection.
func (s *BidirectionalStreamingResultWithExplicitViewMethodServerStream) Send(v *bidirectionalstreamingresultwithexplicitviewservice.Usertype) error {
	var err error
	// Upgrade the HTTP connection to a websocket connection only once. Connection
	// upgrade is done here so that authorization logic in the endpoint is executed
	// before calling the actual service method which may call Send().
	s.once.Do(func() {
		var conn *websocket.Conn
		conn, err = s.upgrader.Upgrade(s.w, s.r, nil)
		if err != nil {
			return
		}
		if s.configurer != nil {
			conn = s.configurer(conn, s.cancel)
		}
		s.conn = conn
	})
	if err != nil {
		return err
	}
	res := bidirectionalstreamingresultwithexplicitviewservice.NewViewedUsertype(v, "extended")
	body := NewBidirectionalStreamingResultWithExplicitViewMethodResponseBodyExtended(res.Projected)
	return s.conn.WriteJSON(body)
}
`

var BidirectionalStreamingResultWithExplicitViewServerStreamRecvCode = `// Recv reads instances of "float32" from the
// "BidirectionalStreamingResultWithExplicitViewMethod" endpoint websocket
// connection.
func (s *BidirectionalStreamingResultWithExplicitViewMethodServerStream) Recv() (float32, error) {
	var (
		rv  float32
		msg *float32
		err error
	)
	// Upgrade the HTTP connection to a websocket connection only once. Connection
	// upgrade is done here so that authorization logic in the endpoint is executed
	// before calling the actual service method which may call Recv().
	s.once.Do(func() {
		var conn *websocket.Conn
		conn, err = s.upgrader.Upgrade(s.w, s.r, nil)
		if err != nil {
			return
		}
		if s.configurer != nil {
			conn = s.configurer(conn, s.cancel)
		}
		s.conn = conn
	})
	if err != nil {
		return rv, err
	}
	if err = s.conn.ReadJSON(&msg); err != nil {
		return rv, err
	}
	if msg == nil {
		return rv, io.EOF
	}
	return *msg, nil
}
`

var BidirectionalStreamingResultWithExplicitViewClientStreamSendCode = `// Send streams instances of "float32" to the
// "BidirectionalStreamingResultWithExplicitViewMethod" endpoint websocket
// connection.
func (s *BidirectionalStreamingResultWithExplicitViewMethodClientStream) Send(v float32) error {
	return s.conn.WriteJSON(v)
}
`

var BidirectionalStreamingResultWithExplicitViewClientStreamRecvCode = `// Recv reads instances of
// "bidirectionalstreamingresultwithexplicitviewservice.Usertype" from the
// "BidirectionalStreamingResultWithExplicitViewMethod" endpoint websocket
// connection.
func (s *BidirectionalStreamingResultWithExplicitViewMethodClientStream) Recv() (*bidirectionalstreamingresultwithexplicitviewservice.Usertype, error) {
	var (
		rv   *bidirectionalstreamingresultwithexplicitviewservice.Usertype
		body BidirectionalStreamingResultWithExplicitViewMethodResponseBody
		err  error
	)
	err = s.conn.ReadJSON(&body)
	if websocket.IsCloseError(err, websocket.CloseNormalClosure) {
		return rv, io.EOF
	}
	if err != nil {
		return rv, err
	}
	res := NewBidirectionalStreamingResultWithExplicitViewMethodUsertypeOK(&body)
	vres := &bidirectionalstreamingresultwithexplicitviewserviceviews.Usertype{res, "extended"}
	if err := bidirectionalstreamingresultwithexplicitviewserviceviews.ValidateUsertype(vres); err != nil {
		return rv, goahttp.ErrValidationError("BidirectionalStreamingResultWithExplicitViewService", "BidirectionalStreamingResultWithExplicitViewMethod", err)
	}
	return bidirectionalstreamingresultwithexplicitviewservice.NewUsertype(vres), nil
}
`

var BidirectionalStreamingResultListWithViewsServerStreamSendCode = `// Send streams instances of
// "bidirectionalstreamingresultlistwithviewsservice.UsertypeList"
// to the "BidirectionalStreamingResultListWithViewsMethod" endpoint
// websocket connection.
func (s *BidirectionalStreamingResultListWithViewsMethodServerStream) Send(v bidirectionalstreamingresultlistwithviewsservice.UsertypeList) error {
	var err error
	// Upgrade the HTTP connection to a websocket connection only once. Connection
	// upgrade is done here so that authorization logic in the endpoint is executed
	// before calling the actual service method which may call Send().
	s.once.Do(func() {
		respHdr := make(http.Header)
		respHdr.Add("goa-view", s.view)
		var conn *websocket.Conn
		conn, err = s.upgrader.Upgrade(s.w, s.r, respHdr)
		if err != nil {
			return
		}
		if s.configurer != nil {
			conn = s.configurer(conn, s.cancel)
		}
		s.conn = conn
	})
	if err != nil {
		return err
	}
	res := bidirectionalstreamingresultlistwithviewsservice.NewViewedUsertypeList(v, s.view)
	var body interface{}
	switch s.view {
	case "tiny":
		body = NewUsertypeResponseTinyList(res.Projected)
	case "extended":
		body = NewUsertypeResponseExtendedList(res.Projected)
	case "default", "":
		body = NewUsertypeResponseList(res.Projected)
	}
	return s.conn.WriteJSON(body)
}
`

var BidirectionalStreamingResultListWithViewsServerStreamRecvCode = `// Recv reads instances of "interface{}" from the
// "BidirectionalStreamingResultListWithViewsMethod" endpoint websocket
// connection.
func (s *BidirectionalStreamingResultListWithViewsMethodServerStream) Recv() (interface{}, error) {
	var (
		rv  interface{}
		msg *interface{}
		err error
	)
	// Upgrade the HTTP connection to a websocket connection only once. Connection
	// upgrade is done here so that authorization logic in the endpoint is executed
	// before calling the actual service method which may call Recv().
	s.once.Do(func() {
		var conn *websocket.Conn
		conn, err = s.upgrader.Upgrade(s.w, s.r, nil)
		if err != nil {
			return
		}
		if s.configurer != nil {
			conn = s.configurer(conn, s.cancel)
		}
		s.conn = conn
	})
	if err != nil {
		return rv, err
	}
	if err = s.conn.ReadJSON(&msg); err != nil {
		return rv, err
	}
	if msg == nil {
		return rv, io.EOF
	}
	return *msg, nil
}
`

var BidirectionalStreamingResultListWithViewsServerStreamSetViewCode = `// SetView sets the view to render the
// bidirectionalstreamingresultlistwithviewsservice.UsertypeList
// type before sending to the
// "BidirectionalStreamingResultListWithViewsMethod" endpoint websocket
// connection.
func (s *BidirectionalStreamingResultListWithViewsMethodServerStream) SetView(view string) {
	s.view = view
}
`

var BidirectionalStreamingResultListWithViewsClientStreamSendCode = `// Send streams instances of "interface{}" to the
// "BidirectionalStreamingResultListWithViewsMethod" endpoint websocket
// connection.
func (s *BidirectionalStreamingResultListWithViewsMethodClientStream) Send(v interface{}) error {
	return s.conn.WriteJSON(v)
}
`

var BidirectionalStreamingResultListWithViewsClientStreamRecvCode = `// Recv reads instances of
// "bidirectionalstreamingresultlistwithviewsservice.UsertypeList"
// from the "BidirectionalStreamingResultListWithViewsMethod" endpoint
// websocket connection.
func (s *BidirectionalStreamingResultListWithViewsMethodClientStream) Recv() (bidirectionalstreamingresultlistwithviewsservice.UsertypeList, error) {
	var (
		rv   bidirectionalstreamingresultlistwithviewsservice.UsertypeList
		body BidirectionalStreamingResultListWithViewsMethodResponseBody
		err  error
	)
	err = s.conn.ReadJSON(&body)
	if websocket.IsCloseError(err, websocket.CloseNormalClosure) {
		return rv, io.EOF
	}
	if err != nil {
		return rv, err
	}
	res := NewBidirectionalStreamingResultListWithViewsMethodUsertypeListOK(body)
	vres := bidirectionalstreamingresultlistwithviewsserviceviews.UsertypeList{res, s.view}
	if err := bidirectionalstreamingresultlistwithviewsserviceviews.ValidateUsertypeList(vres); err != nil {
		return rv, goahttp.ErrValidationError("BidirectionalStreamingResultListWithViewsService", "BidirectionalStreamingResultListWithViewsMethod", err)
	}
	return bidirectionalstreamingresultlistwithviewsservice.NewUsertypeList(vres), nil
}
`

var BidirectionalStreamingResultListWithViewsClientStreamSetViewCode = `// SetView sets the view to render the interface{} type before sending to the
// "BidirectionalStreamingResultListWithViewsMethod" endpoint websocket
// connection.
func (s *BidirectionalStreamingResultListWithViewsMethodClientStream) SetView(view string) {
	s.view = view
}
`

var BidirectionalStreamingResultListWithExplicitViewServerStreamSendCode = `// Send streams instances of
// "bidirectionalstreamingresultlistwithexplicitviewservice.UsertypeList"
// to the "BidirectionalStreamingResultListWithExplicitViewMethod"
// endpoint websocket connection.
func (s *BidirectionalStreamingResultListWithExplicitViewMethodServerStream) Send(v bidirectionalstreamingresultlistwithexplicitviewservice.UsertypeList) error {
	var err error
	// Upgrade the HTTP connection to a websocket connection only once. Connection
	// upgrade is done here so that authorization logic in the endpoint is executed
	// before calling the actual service method which may call Send().
	s.once.Do(func() {
		var conn *websocket.Conn
		conn, err = s.upgrader.Upgrade(s.w, s.r, nil)
		if err != nil {
			return
		}
		if s.configurer != nil {
			conn = s.configurer(conn, s.cancel)
		}
		s.conn = conn
	})
	if err != nil {
		return err
	}
	res := bidirectionalstreamingresultlistwithexplicitviewservice.NewViewedUsertypeList(v, "tiny")
	body := NewUsertypeResponseTinyList(res.Projected)
	return s.conn.WriteJSON(body)
}
`

var BidirectionalStreamingResultListWithExplicitViewServerStreamRecvCode = `// Recv reads instances of "interface{}" from the
// "BidirectionalStreamingResultListWithExplicitViewMethod" endpoint
// websocket connection.
func (s *BidirectionalStreamingResultListWithExplicitViewMethodServerStream) Recv() (interface{}, error) {
	var (
		rv  interface{}
		msg *interface{}
		err error
	)
	// Upgrade the HTTP connection to a websocket connection only once. Connection
	// upgrade is done here so that authorization logic in the endpoint is executed
	// before calling the actual service method which may call Recv().
	s.once.Do(func() {
		var conn *websocket.Conn
		conn, err = s.upgrader.Upgrade(s.w, s.r, nil)
		if err != nil {
			return
		}
		if s.configurer != nil {
			conn = s.configurer(conn, s.cancel)
		}
		s.conn = conn
	})
	if err != nil {
		return rv, err
	}
	if err = s.conn.ReadJSON(&msg); err != nil {
		return rv, err
	}
	if msg == nil {
		return rv, io.EOF
	}
	return *msg, nil
}
`

var BidirectionalStreamingResultListWithExplicitViewClientStreamSendCode = `// Send streams instances of "interface{}" to the
// "BidirectionalStreamingResultListWithExplicitViewMethod" endpoint
// websocket connection.
func (s *BidirectionalStreamingResultListWithExplicitViewMethodClientStream) Send(v interface{}) error {
	return s.conn.WriteJSON(v)
}
`

var BidirectionalStreamingResultListWithExplicitViewClientStreamRecvCode = `// Recv reads instances of
// "bidirectionalstreamingresultlistwithexplicitviewservice.UsertypeList"
// from the "BidirectionalStreamingResultListWithExplicitViewMethod"
// endpoint websocket connection.
func (s *BidirectionalStreamingResultListWithExplicitViewMethodClientStream) Recv() (bidirectionalstreamingresultlistwithexplicitviewservice.UsertypeList, error) {
	var (
		rv   bidirectionalstreamingresultlistwithexplicitviewservice.UsertypeList
		body BidirectionalStreamingResultListWithExplicitViewMethodResponseBody
		err  error
	)
	err = s.conn.ReadJSON(&body)
	if websocket.IsCloseError(err, websocket.CloseNormalClosure) {
		return rv, io.EOF
	}
	if err != nil {
		return rv, err
	}
	res := NewBidirectionalStreamingResultListWithExplicitViewMethodUsertypeListOK(body)
	vres := bidirectionalstreamingresultlistwithexplicitviewserviceviews.UsertypeList{res, "tiny"}
	if err := bidirectionalstreamingresultlistwithexplicitviewserviceviews.ValidateUsertypeList(vres); err != nil {
		return rv, goahttp.ErrValidationError("BidirectionalStreamingResultListWithExplicitViewService", "BidirectionalStreamingResultListWithExplicitViewMethod", err)
	}
	return bidirectionalstreamingresultlistwithexplicitviewservice.NewUsertypeList(vres), nil
}
`

var BidirectionalStreamingPrimitiveServerStreamSendCode = `// Send streams instances of "string" to the
// "BidirectionalStreamingPrimitiveMethod" endpoint websocket connection.
func (s *BidirectionalStreamingPrimitiveMethodServerStream) Send(v string) error {
	var err error
	// Upgrade the HTTP connection to a websocket connection only once. Connection
	// upgrade is done here so that authorization logic in the endpoint is executed
	// before calling the actual service method which may call Send().
	s.once.Do(func() {
		var conn *websocket.Conn
		conn, err = s.upgrader.Upgrade(s.w, s.r, nil)
		if err != nil {
			return
		}
		if s.configurer != nil {
			conn = s.configurer(conn, s.cancel)
		}
		s.conn = conn
	})
	if err != nil {
		return err
	}
	res := v
	return s.conn.WriteJSON(res)
}
`

var BidirectionalStreamingPrimitiveServerStreamRecvCode = `// Recv reads instances of "string" from the
// "BidirectionalStreamingPrimitiveMethod" endpoint websocket connection.
func (s *BidirectionalStreamingPrimitiveMethodServerStream) Recv() (string, error) {
	var (
		rv  string
		msg *string
		err error
	)
	// Upgrade the HTTP connection to a websocket connection only once. Connection
	// upgrade is done here so that authorization logic in the endpoint is executed
	// before calling the actual service method which may call Recv().
	s.once.Do(func() {
		var conn *websocket.Conn
		conn, err = s.upgrader.Upgrade(s.w, s.r, nil)
		if err != nil {
			return
		}
		if s.configurer != nil {
			conn = s.configurer(conn, s.cancel)
		}
		s.conn = conn
	})
	if err != nil {
		return rv, err
	}
	if err = s.conn.ReadJSON(&msg); err != nil {
		return rv, err
	}
	if msg == nil {
		return rv, io.EOF
	}
	return *msg, nil
}
`

var BidirectionalStreamingPrimitiveClientStreamSendCode = `// Send streams instances of "string" to the
// "BidirectionalStreamingPrimitiveMethod" endpoint websocket connection.
func (s *BidirectionalStreamingPrimitiveMethodClientStream) Send(v string) error {
	return s.conn.WriteJSON(v)
}
`

var BidirectionalStreamingPrimitiveClientStreamRecvCode = `// Recv reads instances of "string" from the
// "BidirectionalStreamingPrimitiveMethod" endpoint websocket connection.
func (s *BidirectionalStreamingPrimitiveMethodClientStream) Recv() (string, error) {
	var (
		rv   string
		body string
		err  error
	)
	err = s.conn.ReadJSON(&body)
	if websocket.IsCloseError(err, websocket.CloseNormalClosure) {
		return rv, io.EOF
	}
	if err != nil {
		return rv, err
	}
	return body, nil
}
`

var BidirectionalStreamingPrimitiveArrayServerStreamSendCode = `// Send streams instances of "[]string" to the
// "BidirectionalStreamingPrimitiveArrayMethod" endpoint websocket connection.
func (s *BidirectionalStreamingPrimitiveArrayMethodServerStream) Send(v []string) error {
	var err error
	// Upgrade the HTTP connection to a websocket connection only once. Connection
	// upgrade is done here so that authorization logic in the endpoint is executed
	// before calling the actual service method which may call Send().
	s.once.Do(func() {
		var conn *websocket.Conn
		conn, err = s.upgrader.Upgrade(s.w, s.r, nil)
		if err != nil {
			return
		}
		if s.configurer != nil {
			conn = s.configurer(conn, s.cancel)
		}
		s.conn = conn
	})
	if err != nil {
		return err
	}
	res := v
	return s.conn.WriteJSON(res)
}
`

var BidirectionalStreamingPrimitiveArrayServerStreamRecvCode = `// Recv reads instances of "[]int32" from the
// "BidirectionalStreamingPrimitiveArrayMethod" endpoint websocket connection.
func (s *BidirectionalStreamingPrimitiveArrayMethodServerStream) Recv() ([]int32, error) {
	var (
		rv   []int32
		body []int32
		err  error
	)
	// Upgrade the HTTP connection to a websocket connection only once. Connection
	// upgrade is done here so that authorization logic in the endpoint is executed
	// before calling the actual service method which may call Recv().
	s.once.Do(func() {
		var conn *websocket.Conn
		conn, err = s.upgrader.Upgrade(s.w, s.r, nil)
		if err != nil {
			return
		}
		if s.configurer != nil {
			conn = s.configurer(conn, s.cancel)
		}
		s.conn = conn
	})
	if err != nil {
		return rv, err
	}
	if err = s.conn.ReadJSON(&body); err != nil {
		return rv, err
	}
	if body == nil {
		return rv, io.EOF
	}
	return body, nil
}
`

var BidirectionalStreamingPrimitiveArrayClientStreamSendCode = `// Send streams instances of "[]int32" to the
// "BidirectionalStreamingPrimitiveArrayMethod" endpoint websocket connection.
func (s *BidirectionalStreamingPrimitiveArrayMethodClientStream) Send(v []int32) error {
	return s.conn.WriteJSON(v)
}
`

var BidirectionalStreamingPrimitiveArrayClientStreamRecvCode = `// Recv reads instances of "[]string" from the
// "BidirectionalStreamingPrimitiveArrayMethod" endpoint websocket connection.
func (s *BidirectionalStreamingPrimitiveArrayMethodClientStream) Recv() ([]string, error) {
	var (
		rv   []string
		body []string
		err  error
	)
	err = s.conn.ReadJSON(&body)
	if websocket.IsCloseError(err, websocket.CloseNormalClosure) {
		return rv, io.EOF
	}
	if err != nil {
		return rv, err
	}
	return body, nil
}
`

var BidirectionalStreamingPrimitiveMapServerStreamSendCode = `// Send streams instances of "map[int]int" to the
// "BidirectionalStreamingPrimitiveMapMethod" endpoint websocket connection.
func (s *BidirectionalStreamingPrimitiveMapMethodServerStream) Send(v map[int]int) error {
	var err error
	// Upgrade the HTTP connection to a websocket connection only once. Connection
	// upgrade is done here so that authorization logic in the endpoint is executed
	// before calling the actual service method which may call Send().
	s.once.Do(func() {
		var conn *websocket.Conn
		conn, err = s.upgrader.Upgrade(s.w, s.r, nil)
		if err != nil {
			return
		}
		if s.configurer != nil {
			conn = s.configurer(conn, s.cancel)
		}
		s.conn = conn
	})
	if err != nil {
		return err
	}
	res := v
	return s.conn.WriteJSON(res)
}
`

var BidirectionalStreamingPrimitiveMapServerStreamRecvCode = `// Recv reads instances of "map[string]int32" from the
// "BidirectionalStreamingPrimitiveMapMethod" endpoint websocket connection.
func (s *BidirectionalStreamingPrimitiveMapMethodServerStream) Recv() (map[string]int32, error) {
	var (
		rv   map[string]int32
		body map[string]int32
		err  error
	)
	// Upgrade the HTTP connection to a websocket connection only once. Connection
	// upgrade is done here so that authorization logic in the endpoint is executed
	// before calling the actual service method which may call Recv().
	s.once.Do(func() {
		var conn *websocket.Conn
		conn, err = s.upgrader.Upgrade(s.w, s.r, nil)
		if err != nil {
			return
		}
		if s.configurer != nil {
			conn = s.configurer(conn, s.cancel)
		}
		s.conn = conn
	})
	if err != nil {
		return rv, err
	}
	if err = s.conn.ReadJSON(&body); err != nil {
		return rv, err
	}
	if body == nil {
		return rv, io.EOF
	}
	return body, nil
}
`

var BidirectionalStreamingPrimitiveMapClientStreamSendCode = `// Send streams instances of "map[string]int32" to the
// "BidirectionalStreamingPrimitiveMapMethod" endpoint websocket connection.
func (s *BidirectionalStreamingPrimitiveMapMethodClientStream) Send(v map[string]int32) error {
	return s.conn.WriteJSON(v)
}
`

var BidirectionalStreamingPrimitiveMapClientStreamRecvCode = `// Recv reads instances of "map[int]int" from the
// "BidirectionalStreamingPrimitiveMapMethod" endpoint websocket connection.
func (s *BidirectionalStreamingPrimitiveMapMethodClientStream) Recv() (map[int]int, error) {
	var (
		rv   map[int]int
		body map[int]int
		err  error
	)
	err = s.conn.ReadJSON(&body)
	if websocket.IsCloseError(err, websocket.CloseNormalClosure) {
		return rv, io.EOF
	}
	if err != nil {
		return rv, err
	}
	return body, nil
}
`

var BidirectionalStreamingUserTypeArrayServerStreamSendCode = `// Send streams instances of
// "[]*bidirectionalstreamingusertypearrayservice.ResultType" to the
// "BidirectionalStreamingUserTypeArrayMethod" endpoint websocket connection.
func (s *BidirectionalStreamingUserTypeArrayMethodServerStream) Send(v []*bidirectionalstreamingusertypearrayservice.ResultType) error {
	var err error
	// Upgrade the HTTP connection to a websocket connection only once. Connection
	// upgrade is done here so that authorization logic in the endpoint is executed
	// before calling the actual service method which may call Send().
	s.once.Do(func() {
		var conn *websocket.Conn
		conn, err = s.upgrader.Upgrade(s.w, s.r, nil)
		if err != nil {
			return
		}
		if s.configurer != nil {
			conn = s.configurer(conn, s.cancel)
		}
		s.conn = conn
	})
	if err != nil {
		return err
	}
	res := v
	body := NewBidirectionalStreamingUserTypeArrayMethodResponseBody(res)
	return s.conn.WriteJSON(body)
}
`

var BidirectionalStreamingUserTypeArrayServerStreamRecvCode = `// Recv reads instances of
// "[]*bidirectionalstreamingusertypearrayservice.RequestType" from the
// "BidirectionalStreamingUserTypeArrayMethod" endpoint websocket connection.
func (s *BidirectionalStreamingUserTypeArrayMethodServerStream) Recv() ([]*bidirectionalstreamingusertypearrayservice.RequestType, error) {
	var (
		rv   []*bidirectionalstreamingusertypearrayservice.RequestType
		body []*RequestType
		err  error
	)
	// Upgrade the HTTP connection to a websocket connection only once. Connection
	// upgrade is done here so that authorization logic in the endpoint is executed
	// before calling the actual service method which may call Recv().
	s.once.Do(func() {
		var conn *websocket.Conn
		conn, err = s.upgrader.Upgrade(s.w, s.r, nil)
		if err != nil {
			return
		}
		if s.configurer != nil {
			conn = s.configurer(conn, s.cancel)
		}
		s.conn = conn
	})
	if err != nil {
		return rv, err
	}
	if err = s.conn.ReadJSON(&body); err != nil {
		return rv, err
	}
	if body == nil {
		return rv, io.EOF
	}
	return NewBidirectionalStreamingUserTypeArrayMethodArray(body), nil
}
`

var BidirectionalStreamingUserTypeArrayClientStreamSendCode = `// Send streams instances of
// "[]*bidirectionalstreamingusertypearrayservice.RequestType" to the
// "BidirectionalStreamingUserTypeArrayMethod" endpoint websocket connection.
func (s *BidirectionalStreamingUserTypeArrayMethodClientStream) Send(v []*bidirectionalstreamingusertypearrayservice.RequestType) error {
	body := NewRequestType(v)
	return s.conn.WriteJSON(body)
}
`

var BidirectionalStreamingUserTypeArrayClientStreamRecvCode = `// Recv reads instances of
// "[]*bidirectionalstreamingusertypearrayservice.ResultType" from the
// "BidirectionalStreamingUserTypeArrayMethod" endpoint websocket connection.
func (s *BidirectionalStreamingUserTypeArrayMethodClientStream) Recv() ([]*bidirectionalstreamingusertypearrayservice.ResultType, error) {
	var (
		rv   []*bidirectionalstreamingusertypearrayservice.ResultType
		body BidirectionalStreamingUserTypeArrayMethodResponseBody
		err  error
	)
	err = s.conn.ReadJSON(&body)
	if websocket.IsCloseError(err, websocket.CloseNormalClosure) {
		return rv, io.EOF
	}
	if err != nil {
		return rv, err
	}
	res := NewBidirectionalStreamingUserTypeArrayMethodResultTypeOK(body)
	return res, nil
}
`

var BidirectionalStreamingUserTypeMapServerStreamSendCode = `// Send streams instances of
// "map[string]*bidirectionalstreamingusertypemapservice.ResultType" to the
// "BidirectionalStreamingUserTypeMapMethod" endpoint websocket connection.
func (s *BidirectionalStreamingUserTypeMapMethodServerStream) Send(v map[string]*bidirectionalstreamingusertypemapservice.ResultType) error {
	var err error
	// Upgrade the HTTP connection to a websocket connection only once. Connection
	// upgrade is done here so that authorization logic in the endpoint is executed
	// before calling the actual service method which may call Send().
	s.once.Do(func() {
		var conn *websocket.Conn
		conn, err = s.upgrader.Upgrade(s.w, s.r, nil)
		if err != nil {
			return
		}
		if s.configurer != nil {
			conn = s.configurer(conn, s.cancel)
		}
		s.conn = conn
	})
	if err != nil {
		return err
	}
	res := v
	body := NewBidirectionalStreamingUserTypeMapMethodResponseBody(res)
	return s.conn.WriteJSON(body)
}
`

var BidirectionalStreamingUserTypeMapServerStreamRecvCode = `// Recv reads instances of
// "map[string]*bidirectionalstreamingusertypemapservice.RequestType" from the
// "BidirectionalStreamingUserTypeMapMethod" endpoint websocket connection.
func (s *BidirectionalStreamingUserTypeMapMethodServerStream) Recv() (map[string]*bidirectionalstreamingusertypemapservice.RequestType, error) {
	var (
		rv   map[string]*bidirectionalstreamingusertypemapservice.RequestType
		body map[string]*RequestType
		err  error
	)
	// Upgrade the HTTP connection to a websocket connection only once. Connection
	// upgrade is done here so that authorization logic in the endpoint is executed
	// before calling the actual service method which may call Recv().
	s.once.Do(func() {
		var conn *websocket.Conn
		conn, err = s.upgrader.Upgrade(s.w, s.r, nil)
		if err != nil {
			return
		}
		if s.configurer != nil {
			conn = s.configurer(conn, s.cancel)
		}
		s.conn = conn
	})
	if err != nil {
		return rv, err
	}
	if err = s.conn.ReadJSON(&body); err != nil {
		return rv, err
	}
	if body == nil {
		return rv, io.EOF
	}
	return NewBidirectionalStreamingUserTypeMapMethodMap(body), nil
}
`

var BidirectionalStreamingUserTypeMapClientStreamSendCode = `// Send streams instances of
// "map[string]*bidirectionalstreamingusertypemapservice.RequestType" to the
// "BidirectionalStreamingUserTypeMapMethod" endpoint websocket connection.
func (s *BidirectionalStreamingUserTypeMapMethodClientStream) Send(v map[string]*bidirectionalstreamingusertypemapservice.RequestType) error {
	body := NewMapStringRequestType(v)
	return s.conn.WriteJSON(body)
}
`

var BidirectionalStreamingUserTypeMapClientStreamRecvCode = `// Recv reads instances of
// "map[string]*bidirectionalstreamingusertypemapservice.ResultType" from the
// "BidirectionalStreamingUserTypeMapMethod" endpoint websocket connection.
func (s *BidirectionalStreamingUserTypeMapMethodClientStream) Recv() (map[string]*bidirectionalstreamingusertypemapservice.ResultType, error) {
	var (
		rv   map[string]*bidirectionalstreamingusertypemapservice.ResultType
		body BidirectionalStreamingUserTypeMapMethodResponseBody
		err  error
	)
	err = s.conn.ReadJSON(&body)
	if websocket.IsCloseError(err, websocket.CloseNormalClosure) {
		return rv, io.EOF
	}
	if err != nil {
		return rv, err
	}
	res := NewBidirectionalStreamingUserTypeMapMethodMapStringResultTypeOK(body)
	return res, nil
}
`
