// DO NOT EDIT!
// Code generated by ffjson <https://github.com/pquerna/ffjson>
// source: consume_stats.go
// DO NOT EDIT!

package rocketmq

import (
	"bytes"
	"fmt"

	fflib "github.com/pquerna/ffjson/fflib/v1"
)

func (mj *ConsumeStats) MarshalJSON() ([]byte, error) {
	var buf fflib.Buffer
	if mj == nil {
		buf.WriteString("null")
		return buf.Bytes(), nil
	}
	err := mj.MarshalJSONBuf(&buf)
	if err != nil {
		return nil, err
	}
	return buf.Bytes(), nil
}
func (mj *ConsumeStats) MarshalJSONBuf(buf fflib.EncodingBuffer) error {
	if mj == nil {
		buf.WriteString("null")
		return nil
	}
	var err error
	var obj []byte
	_ = obj
	_ = err
	/* Falling back. type=map[rocketmq.MessageQueue]rocketmq.OffsetWrapper kind=map */
	buf.WriteString(`{"offsetTable":`)
	err = buf.Encode(mj.OffsetTable)
	if err != nil {
		return err
	}
	buf.WriteString(`,"consumeTps":`)
	fflib.FormatBits2(buf, uint64(mj.ConsumeTps), 10, mj.ConsumeTps < 0)
	buf.WriteByte('}')
	return nil
}

const (
	ffj_t_ConsumeStatsbase = iota
	ffj_t_ConsumeStatsno_such_key

	ffj_t_ConsumeStats_OffsetTable

	ffj_t_ConsumeStats_ConsumeTps
)

var ffj_key_ConsumeStats_OffsetTable = []byte("offsetTable")

var ffj_key_ConsumeStats_ConsumeTps = []byte("consumeTps")

func (uj *ConsumeStats) UnmarshalJSON(input []byte) error {
	fs := fflib.NewFFLexer(input)
	return uj.UnmarshalJSONFFLexer(fs, fflib.FFParse_map_start)
}

func (uj *ConsumeStats) UnmarshalJSONFFLexer(fs *fflib.FFLexer, state fflib.FFParseState) error {
	var err error = nil
	currentKey := ffj_t_ConsumeStatsbase
	_ = currentKey
	tok := fflib.FFTok_init
	wantedTok := fflib.FFTok_init

mainparse:
	for {
		tok = fs.Scan()
		//	println(fmt.Sprintf("debug: tok: %v  state: %v", tok, state))
		if tok == fflib.FFTok_error {
			goto tokerror
		}

		switch state {

		case fflib.FFParse_map_start:
			if tok != fflib.FFTok_left_bracket {
				wantedTok = fflib.FFTok_left_bracket
				goto wrongtokenerror
			}
			state = fflib.FFParse_want_key
			continue

		case fflib.FFParse_after_value:
			if tok == fflib.FFTok_comma {
				state = fflib.FFParse_want_key
			} else if tok == fflib.FFTok_right_bracket {
				goto done
			} else {
				wantedTok = fflib.FFTok_comma
				goto wrongtokenerror
			}

		case fflib.FFParse_want_key:
			// json {} ended. goto exit. woo.
			if tok == fflib.FFTok_right_bracket {
				goto done
			}
			if tok != fflib.FFTok_string {
				wantedTok = fflib.FFTok_string
				goto wrongtokenerror
			}

			kn := fs.Output.Bytes()
			if len(kn) <= 0 {
				// "" case. hrm.
				currentKey = ffj_t_ConsumeStatsno_such_key
				state = fflib.FFParse_want_colon
				goto mainparse
			} else {
				switch kn[0] {

				case 'c':

					if bytes.Equal(ffj_key_ConsumeStats_ConsumeTps, kn) {
						currentKey = ffj_t_ConsumeStats_ConsumeTps
						state = fflib.FFParse_want_colon
						goto mainparse
					}

				case 'o':

					if bytes.Equal(ffj_key_ConsumeStats_OffsetTable, kn) {
						currentKey = ffj_t_ConsumeStats_OffsetTable
						state = fflib.FFParse_want_colon
						goto mainparse
					}

				}

				if fflib.EqualFoldRight(ffj_key_ConsumeStats_ConsumeTps, kn) {
					currentKey = ffj_t_ConsumeStats_ConsumeTps
					state = fflib.FFParse_want_colon
					goto mainparse
				}

				if fflib.EqualFoldRight(ffj_key_ConsumeStats_OffsetTable, kn) {
					currentKey = ffj_t_ConsumeStats_OffsetTable
					state = fflib.FFParse_want_colon
					goto mainparse
				}

				currentKey = ffj_t_ConsumeStatsno_such_key
				state = fflib.FFParse_want_colon
				goto mainparse
			}

		case fflib.FFParse_want_colon:
			if tok != fflib.FFTok_colon {
				wantedTok = fflib.FFTok_colon
				goto wrongtokenerror
			}
			state = fflib.FFParse_want_value
			continue
		case fflib.FFParse_want_value:

			if tok == fflib.FFTok_left_brace || tok == fflib.FFTok_left_bracket || tok == fflib.FFTok_integer || tok == fflib.FFTok_double || tok == fflib.FFTok_string || tok == fflib.FFTok_bool || tok == fflib.FFTok_null {
				switch currentKey {

				case ffj_t_ConsumeStats_OffsetTable:
					goto handle_OffsetTable

				case ffj_t_ConsumeStats_ConsumeTps:
					goto handle_ConsumeTps

				case ffj_t_ConsumeStatsno_such_key:
					err = fs.SkipField(tok)
					if err != nil {
						return fs.WrapErr(err)
					}
					state = fflib.FFParse_after_value
					goto mainparse
				}
			} else {
				goto wantedvalue
			}
		}
	}

handle_OffsetTable:

	/* handler: uj.OffsetTable type=map[rocketmq.MessageQueue]rocketmq.OffsetWrapper kind=map quoted=false*/

	{

		{
			if tok != fflib.FFTok_left_bracket && tok != fflib.FFTok_null {
				return fs.WrapErr(fmt.Errorf("cannot unmarshal %s into Go value for ", tok))
			}
		}

		if tok == fflib.FFTok_null {
			uj.OffsetTable = nil
		} else {

			uj.OffsetTable = make(map[MessageQueue]OffsetWrapper, 0)

			wantVal := true

			for {

				var k MessageQueue

				var tmp_uj__OffsetTable OffsetWrapper

				tok = fs.Scan()
				if tok == fflib.FFTok_error {
					goto tokerror
				}
				if tok == fflib.FFTok_right_bracket {
					break
				}

				if tok == fflib.FFTok_comma {
					if wantVal == true {
						// TODO(pquerna): this isn't an ideal error message, this handles
						// things like [,,,] as an array value.
						return fs.WrapErr(fmt.Errorf("wanted value token, but got token: %v", tok))
					}
					continue
				} else {
					wantVal = true
				}

				/* handler: k type=rocketmq.MessageQueue kind=struct quoted=false*/

				{
					if tok == fflib.FFTok_null {

						state = fflib.FFParse_after_value
						goto mainparse
					}

					err = k.UnmarshalJSONFFLexer(fs, fflib.FFParse_want_key)
					if err != nil {
						return err
					}
					state = fflib.FFParse_after_value
				}

				// Expect ':' after key
				tok = fs.Scan()
				if tok != fflib.FFTok_colon {
					return fs.WrapErr(fmt.Errorf("wanted colon token, but got token: %v", tok))
				}

				tok = fs.Scan()
				/* handler: tmp_uj__OffsetTable type=rocketmq.OffsetWrapper kind=struct quoted=false*/

				{
					if tok == fflib.FFTok_null {

						state = fflib.FFParse_after_value
						goto mainparse
					}

					err = tmp_uj__OffsetTable.UnmarshalJSONFFLexer(fs, fflib.FFParse_want_key)
					if err != nil {
						return err
					}
					state = fflib.FFParse_after_value
				}

				uj.OffsetTable[k] = tmp_uj__OffsetTable

				wantVal = false
			}

		}
	}

	state = fflib.FFParse_after_value
	goto mainparse

handle_ConsumeTps:

	/* handler: uj.ConsumeTps type=int64 kind=int64 quoted=false*/

	{
		if tok != fflib.FFTok_integer && tok != fflib.FFTok_null {
			return fs.WrapErr(fmt.Errorf("cannot unmarshal %s into Go value for int64", tok))
		}
	}

	{

		if tok == fflib.FFTok_null {

		} else {

			tval, err := fflib.ParseInt(fs.Output.Bytes(), 10, 64)

			if err != nil {
				return fs.WrapErr(err)
			}

			uj.ConsumeTps = int64(tval)

		}
	}

	state = fflib.FFParse_after_value
	goto mainparse

wantedvalue:
	return fs.WrapErr(fmt.Errorf("wanted value token, but got token: %v", tok))
wrongtokenerror:
	return fs.WrapErr(fmt.Errorf("ffjson: wanted token: %v, but got token: %v output=%s", wantedTok, tok, fs.Output.String()))
tokerror:
	if fs.BigError != nil {
		return fs.WrapErr(fs.BigError)
	}
	err = fs.Error.ToError()
	if err != nil {
		return fs.WrapErr(err)
	}
	panic("ffjson-generated: unreachable, please report bug.")
done:

	return nil
}

func (mj *MessageQueue) MarshalJSON() ([]byte, error) {
	var buf fflib.Buffer
	if mj == nil {
		buf.WriteString("null")
		return buf.Bytes(), nil
	}
	err := mj.MarshalJSONBuf(&buf)
	if err != nil {
		return nil, err
	}
	return buf.Bytes(), nil
}
func (mj *MessageQueue) MarshalJSONBuf(buf fflib.EncodingBuffer) error {
	if mj == nil {
		buf.WriteString("null")
		return nil
	}
	var err error
	var obj []byte
	_ = obj
	_ = err
	buf.WriteString(`{"topic":`)
	fflib.WriteJsonString(buf, string(mj.Topic))
	buf.WriteString(`,"brokerName":`)
	fflib.WriteJsonString(buf, string(mj.BrokerName))
	buf.WriteString(`,"queueId":`)
	fflib.FormatBits2(buf, uint64(mj.QueueId), 10, mj.QueueId < 0)
	buf.WriteByte('}')
	return nil
}

const (
	ffj_t_MessageQueuebase = iota
	ffj_t_MessageQueueno_such_key

	ffj_t_MessageQueue_Topic

	ffj_t_MessageQueue_BrokerName

	ffj_t_MessageQueue_QueueId
)

var ffj_key_MessageQueue_Topic = []byte("topic")

var ffj_key_MessageQueue_BrokerName = []byte("brokerName")

var ffj_key_MessageQueue_QueueId = []byte("queueId")

func (uj *MessageQueue) UnmarshalJSON(input []byte) error {
	fs := fflib.NewFFLexer(input)
	return uj.UnmarshalJSONFFLexer(fs, fflib.FFParse_map_start)
}

func (uj *MessageQueue) UnmarshalJSONFFLexer(fs *fflib.FFLexer, state fflib.FFParseState) error {
	var err error = nil
	currentKey := ffj_t_MessageQueuebase
	_ = currentKey
	tok := fflib.FFTok_init
	wantedTok := fflib.FFTok_init

mainparse:
	for {
		tok = fs.Scan()
		//	println(fmt.Sprintf("debug: tok: %v  state: %v", tok, state))
		if tok == fflib.FFTok_error {
			goto tokerror
		}

		switch state {

		case fflib.FFParse_map_start:
			if tok != fflib.FFTok_left_bracket {
				wantedTok = fflib.FFTok_left_bracket
				goto wrongtokenerror
			}
			state = fflib.FFParse_want_key
			continue

		case fflib.FFParse_after_value:
			if tok == fflib.FFTok_comma {
				state = fflib.FFParse_want_key
			} else if tok == fflib.FFTok_right_bracket {
				goto done
			} else {
				wantedTok = fflib.FFTok_comma
				goto wrongtokenerror
			}

		case fflib.FFParse_want_key:
			// json {} ended. goto exit. woo.
			if tok == fflib.FFTok_right_bracket {
				goto done
			}
			if tok != fflib.FFTok_string {
				wantedTok = fflib.FFTok_string
				goto wrongtokenerror
			}

			kn := fs.Output.Bytes()
			if len(kn) <= 0 {
				// "" case. hrm.
				currentKey = ffj_t_MessageQueueno_such_key
				state = fflib.FFParse_want_colon
				goto mainparse
			} else {
				switch kn[0] {

				case 'b':

					if bytes.Equal(ffj_key_MessageQueue_BrokerName, kn) {
						currentKey = ffj_t_MessageQueue_BrokerName
						state = fflib.FFParse_want_colon
						goto mainparse
					}

				case 'q':

					if bytes.Equal(ffj_key_MessageQueue_QueueId, kn) {
						currentKey = ffj_t_MessageQueue_QueueId
						state = fflib.FFParse_want_colon
						goto mainparse
					}

				case 't':

					if bytes.Equal(ffj_key_MessageQueue_Topic, kn) {
						currentKey = ffj_t_MessageQueue_Topic
						state = fflib.FFParse_want_colon
						goto mainparse
					}

				}

				if fflib.SimpleLetterEqualFold(ffj_key_MessageQueue_QueueId, kn) {
					currentKey = ffj_t_MessageQueue_QueueId
					state = fflib.FFParse_want_colon
					goto mainparse
				}

				if fflib.EqualFoldRight(ffj_key_MessageQueue_BrokerName, kn) {
					currentKey = ffj_t_MessageQueue_BrokerName
					state = fflib.FFParse_want_colon
					goto mainparse
				}

				if fflib.SimpleLetterEqualFold(ffj_key_MessageQueue_Topic, kn) {
					currentKey = ffj_t_MessageQueue_Topic
					state = fflib.FFParse_want_colon
					goto mainparse
				}

				currentKey = ffj_t_MessageQueueno_such_key
				state = fflib.FFParse_want_colon
				goto mainparse
			}

		case fflib.FFParse_want_colon:
			if tok != fflib.FFTok_colon {
				wantedTok = fflib.FFTok_colon
				goto wrongtokenerror
			}
			state = fflib.FFParse_want_value
			continue
		case fflib.FFParse_want_value:

			if tok == fflib.FFTok_left_brace || tok == fflib.FFTok_left_bracket || tok == fflib.FFTok_integer || tok == fflib.FFTok_double || tok == fflib.FFTok_string || tok == fflib.FFTok_bool || tok == fflib.FFTok_null {
				switch currentKey {

				case ffj_t_MessageQueue_Topic:
					goto handle_Topic

				case ffj_t_MessageQueue_BrokerName:
					goto handle_BrokerName

				case ffj_t_MessageQueue_QueueId:
					goto handle_QueueId

				case ffj_t_MessageQueueno_such_key:
					err = fs.SkipField(tok)
					if err != nil {
						return fs.WrapErr(err)
					}
					state = fflib.FFParse_after_value
					goto mainparse
				}
			} else {
				goto wantedvalue
			}
		}
	}

handle_Topic:

	/* handler: uj.Topic type=string kind=string quoted=false*/

	{

		{
			if tok != fflib.FFTok_string && tok != fflib.FFTok_null {
				return fs.WrapErr(fmt.Errorf("cannot unmarshal %s into Go value for string", tok))
			}
		}

		if tok == fflib.FFTok_null {

		} else {

			outBuf := fs.Output.Bytes()

			uj.Topic = string(string(outBuf))

		}
	}

	state = fflib.FFParse_after_value
	goto mainparse

handle_BrokerName:

	/* handler: uj.BrokerName type=string kind=string quoted=false*/

	{

		{
			if tok != fflib.FFTok_string && tok != fflib.FFTok_null {
				return fs.WrapErr(fmt.Errorf("cannot unmarshal %s into Go value for string", tok))
			}
		}

		if tok == fflib.FFTok_null {

		} else {

			outBuf := fs.Output.Bytes()

			uj.BrokerName = string(string(outBuf))

		}
	}

	state = fflib.FFParse_after_value
	goto mainparse

handle_QueueId:

	/* handler: uj.QueueId type=int32 kind=int32 quoted=false*/

	{
		if tok != fflib.FFTok_integer && tok != fflib.FFTok_null {
			return fs.WrapErr(fmt.Errorf("cannot unmarshal %s into Go value for int32", tok))
		}
	}

	{

		if tok == fflib.FFTok_null {

		} else {

			tval, err := fflib.ParseInt(fs.Output.Bytes(), 10, 32)

			if err != nil {
				return fs.WrapErr(err)
			}

			uj.QueueId = int32(tval)

		}
	}

	state = fflib.FFParse_after_value
	goto mainparse

wantedvalue:
	return fs.WrapErr(fmt.Errorf("wanted value token, but got token: %v", tok))
wrongtokenerror:
	return fs.WrapErr(fmt.Errorf("ffjson: wanted token: %v, but got token: %v output=%s", wantedTok, tok, fs.Output.String()))
tokerror:
	if fs.BigError != nil {
		return fs.WrapErr(fs.BigError)
	}
	err = fs.Error.ToError()
	if err != nil {
		return fs.WrapErr(err)
	}
	panic("ffjson-generated: unreachable, please report bug.")
done:

	return nil
}

func (mj *OffsetWrapper) MarshalJSON() ([]byte, error) {
	var buf fflib.Buffer
	if mj == nil {
		buf.WriteString("null")
		return buf.Bytes(), nil
	}
	err := mj.MarshalJSONBuf(&buf)
	if err != nil {
		return nil, err
	}
	return buf.Bytes(), nil
}
func (mj *OffsetWrapper) MarshalJSONBuf(buf fflib.EncodingBuffer) error {
	if mj == nil {
		buf.WriteString("null")
		return nil
	}
	var err error
	var obj []byte
	_ = obj
	_ = err
	buf.WriteString(`{"brokerOffset":`)
	fflib.FormatBits2(buf, uint64(mj.BrokerOffset), 10, mj.BrokerOffset < 0)
	buf.WriteString(`,"consumerOffset":`)
	fflib.FormatBits2(buf, uint64(mj.ConsumerOffset), 10, mj.ConsumerOffset < 0)
	buf.WriteString(`,"lastTimestamp":`)
	fflib.FormatBits2(buf, uint64(mj.LastTimestamp), 10, mj.LastTimestamp < 0)
	buf.WriteByte('}')
	return nil
}

const (
	ffj_t_OffsetWrapperbase = iota
	ffj_t_OffsetWrapperno_such_key

	ffj_t_OffsetWrapper_BrokerOffset

	ffj_t_OffsetWrapper_ConsumerOffset

	ffj_t_OffsetWrapper_LastTimestamp
)

var ffj_key_OffsetWrapper_BrokerOffset = []byte("brokerOffset")

var ffj_key_OffsetWrapper_ConsumerOffset = []byte("consumerOffset")

var ffj_key_OffsetWrapper_LastTimestamp = []byte("lastTimestamp")

func (uj *OffsetWrapper) UnmarshalJSON(input []byte) error {
	fs := fflib.NewFFLexer(input)
	return uj.UnmarshalJSONFFLexer(fs, fflib.FFParse_map_start)
}

func (uj *OffsetWrapper) UnmarshalJSONFFLexer(fs *fflib.FFLexer, state fflib.FFParseState) error {
	var err error = nil
	currentKey := ffj_t_OffsetWrapperbase
	_ = currentKey
	tok := fflib.FFTok_init
	wantedTok := fflib.FFTok_init

mainparse:
	for {
		tok = fs.Scan()
		//	println(fmt.Sprintf("debug: tok: %v  state: %v", tok, state))
		if tok == fflib.FFTok_error {
			goto tokerror
		}

		switch state {

		case fflib.FFParse_map_start:
			if tok != fflib.FFTok_left_bracket {
				wantedTok = fflib.FFTok_left_bracket
				goto wrongtokenerror
			}
			state = fflib.FFParse_want_key
			continue

		case fflib.FFParse_after_value:
			if tok == fflib.FFTok_comma {
				state = fflib.FFParse_want_key
			} else if tok == fflib.FFTok_right_bracket {
				goto done
			} else {
				wantedTok = fflib.FFTok_comma
				goto wrongtokenerror
			}

		case fflib.FFParse_want_key:
			// json {} ended. goto exit. woo.
			if tok == fflib.FFTok_right_bracket {
				goto done
			}
			if tok != fflib.FFTok_string {
				wantedTok = fflib.FFTok_string
				goto wrongtokenerror
			}

			kn := fs.Output.Bytes()
			if len(kn) <= 0 {
				// "" case. hrm.
				currentKey = ffj_t_OffsetWrapperno_such_key
				state = fflib.FFParse_want_colon
				goto mainparse
			} else {
				switch kn[0] {

				case 'b':

					if bytes.Equal(ffj_key_OffsetWrapper_BrokerOffset, kn) {
						currentKey = ffj_t_OffsetWrapper_BrokerOffset
						state = fflib.FFParse_want_colon
						goto mainparse
					}

				case 'c':

					if bytes.Equal(ffj_key_OffsetWrapper_ConsumerOffset, kn) {
						currentKey = ffj_t_OffsetWrapper_ConsumerOffset
						state = fflib.FFParse_want_colon
						goto mainparse
					}

				case 'l':

					if bytes.Equal(ffj_key_OffsetWrapper_LastTimestamp, kn) {
						currentKey = ffj_t_OffsetWrapper_LastTimestamp
						state = fflib.FFParse_want_colon
						goto mainparse
					}

				}

				if fflib.EqualFoldRight(ffj_key_OffsetWrapper_LastTimestamp, kn) {
					currentKey = ffj_t_OffsetWrapper_LastTimestamp
					state = fflib.FFParse_want_colon
					goto mainparse
				}

				if fflib.EqualFoldRight(ffj_key_OffsetWrapper_ConsumerOffset, kn) {
					currentKey = ffj_t_OffsetWrapper_ConsumerOffset
					state = fflib.FFParse_want_colon
					goto mainparse
				}

				if fflib.EqualFoldRight(ffj_key_OffsetWrapper_BrokerOffset, kn) {
					currentKey = ffj_t_OffsetWrapper_BrokerOffset
					state = fflib.FFParse_want_colon
					goto mainparse
				}

				currentKey = ffj_t_OffsetWrapperno_such_key
				state = fflib.FFParse_want_colon
				goto mainparse
			}

		case fflib.FFParse_want_colon:
			if tok != fflib.FFTok_colon {
				wantedTok = fflib.FFTok_colon
				goto wrongtokenerror
			}
			state = fflib.FFParse_want_value
			continue
		case fflib.FFParse_want_value:

			if tok == fflib.FFTok_left_brace || tok == fflib.FFTok_left_bracket || tok == fflib.FFTok_integer || tok == fflib.FFTok_double || tok == fflib.FFTok_string || tok == fflib.FFTok_bool || tok == fflib.FFTok_null {
				switch currentKey {

				case ffj_t_OffsetWrapper_BrokerOffset:
					goto handle_BrokerOffset

				case ffj_t_OffsetWrapper_ConsumerOffset:
					goto handle_ConsumerOffset

				case ffj_t_OffsetWrapper_LastTimestamp:
					goto handle_LastTimestamp

				case ffj_t_OffsetWrapperno_such_key:
					err = fs.SkipField(tok)
					if err != nil {
						return fs.WrapErr(err)
					}
					state = fflib.FFParse_after_value
					goto mainparse
				}
			} else {
				goto wantedvalue
			}
		}
	}

handle_BrokerOffset:

	/* handler: uj.BrokerOffset type=int64 kind=int64 quoted=false*/

	{
		if tok != fflib.FFTok_integer && tok != fflib.FFTok_null {
			return fs.WrapErr(fmt.Errorf("cannot unmarshal %s into Go value for int64", tok))
		}
	}

	{

		if tok == fflib.FFTok_null {

		} else {

			tval, err := fflib.ParseInt(fs.Output.Bytes(), 10, 64)

			if err != nil {
				return fs.WrapErr(err)
			}

			uj.BrokerOffset = int64(tval)

		}
	}

	state = fflib.FFParse_after_value
	goto mainparse

handle_ConsumerOffset:

	/* handler: uj.ConsumerOffset type=int64 kind=int64 quoted=false*/

	{
		if tok != fflib.FFTok_integer && tok != fflib.FFTok_null {
			return fs.WrapErr(fmt.Errorf("cannot unmarshal %s into Go value for int64", tok))
		}
	}

	{

		if tok == fflib.FFTok_null {

		} else {

			tval, err := fflib.ParseInt(fs.Output.Bytes(), 10, 64)

			if err != nil {
				return fs.WrapErr(err)
			}

			uj.ConsumerOffset = int64(tval)

		}
	}

	state = fflib.FFParse_after_value
	goto mainparse

handle_LastTimestamp:

	/* handler: uj.LastTimestamp type=int64 kind=int64 quoted=false*/

	{
		if tok != fflib.FFTok_integer && tok != fflib.FFTok_null {
			return fs.WrapErr(fmt.Errorf("cannot unmarshal %s into Go value for int64", tok))
		}
	}

	{

		if tok == fflib.FFTok_null {

		} else {

			tval, err := fflib.ParseInt(fs.Output.Bytes(), 10, 64)

			if err != nil {
				return fs.WrapErr(err)
			}

			uj.LastTimestamp = int64(tval)

		}
	}

	state = fflib.FFParse_after_value
	goto mainparse

wantedvalue:
	return fs.WrapErr(fmt.Errorf("wanted value token, but got token: %v", tok))
wrongtokenerror:
	return fs.WrapErr(fmt.Errorf("ffjson: wanted token: %v, but got token: %v output=%s", wantedTok, tok, fs.Output.String()))
tokerror:
	if fs.BigError != nil {
		return fs.WrapErr(fs.BigError)
	}
	err = fs.Error.ToError()
	if err != nil {
		return fs.WrapErr(err)
	}
	panic("ffjson-generated: unreachable, please report bug.")
done:

	return nil
}