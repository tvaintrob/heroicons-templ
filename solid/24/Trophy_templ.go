// Code generated by templ - DO NOT EDIT.

// templ: version: 0.2.476
package solid24

//lint:file-ignore SA4006 This context is only used if a nested component is present.

import "github.com/a-h/templ"
import "context"
import "io"
import "bytes"

func TrophyIcon() templ.Component {
	return templ.ComponentFunc(func(ctx context.Context, templ_7745c5c3_W io.Writer) (templ_7745c5c3_Err error) {
		templ_7745c5c3_Buffer, templ_7745c5c3_IsBuffer := templ_7745c5c3_W.(*bytes.Buffer)
		if !templ_7745c5c3_IsBuffer {
			templ_7745c5c3_Buffer = templ.GetBuffer()
			defer templ.ReleaseBuffer(templ_7745c5c3_Buffer)
		}
		ctx = templ.InitializeContext(ctx)
		templ_7745c5c3_Var1 := templ.GetChildren(ctx)
		if templ_7745c5c3_Var1 == nil {
			templ_7745c5c3_Var1 = templ.NopComponent
		}
		ctx = templ.ClearChildren(ctx)
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<svg width=\"24\" height=\"24\" viewBox=\"0 0 24 24\" fill=\"none\" xmlns=\"http://www.w3.org/2000/svg\" class=\"w-6 h-6\"><path fill-rule=\"evenodd\" clip-rule=\"evenodd\" d=\"M5.16613 2.62136V3.4786C4.13126 3.62665 3.10716 3.80823 2.09497 4.02223C1.70145 4.10543 1.44361 4.48425 1.51057 4.88086C2.02102 7.90456 4.54015 10.2449 7.64915 10.4805C8.4348 11.1221 9.36855 11.5907 10.3916 11.827C10.2927 12.9952 9.8965 14.0776 9.2788 15H8.54088C7.50534 15 6.66588 15.8395 6.66588 16.875V19.5H5.91577C4.67313 19.5 3.66577 20.5074 3.66577 21.75C3.66577 22.1642 4.00156 22.5 4.41577 22.5H19.4158C19.83 22.5 20.1658 22.1642 20.1658 21.75C20.1658 20.5074 19.1584 19.5 17.9158 19.5H17.1659V16.875C17.1659 15.8395 16.3264 15 15.2909 15H14.5524C13.9349 14.0777 13.5389 12.9953 13.4402 11.8271C14.4634 11.5908 15.3973 11.1222 16.1831 10.4805C19.2921 10.2449 21.8113 7.90456 22.3217 4.88086C22.3887 4.48425 22.1308 4.10543 21.7373 4.02223C20.7251 3.80823 19.701 3.62665 18.6661 3.4786V2.62136C18.6661 2.24303 18.3844 1.92394 18.0089 1.87713C16.0127 1.62819 13.9792 1.5 11.9161 1.5C9.85307 1.5 7.81961 1.62819 5.82333 1.87713C5.44791 1.92394 5.16613 2.24303 5.16613 2.62136ZM5.16613 5.25C5.16613 6.44618 5.47762 7.56995 6.02338 8.54424C4.66626 7.9367 3.61376 6.76975 3.16003 5.33687C3.8237 5.20825 4.49252 5.09398 5.16613 4.99438V5.25ZM18.6661 5.25V4.99438C19.3398 5.09398 20.0086 5.20825 20.6722 5.33688C20.2185 6.76975 19.166 7.9367 17.8089 8.54424C18.3547 7.56995 18.6661 6.44618 18.6661 5.25Z\" fill=\"#0F172A\"></path></svg>")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		if !templ_7745c5c3_IsBuffer {
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteTo(templ_7745c5c3_W)
		}
		return templ_7745c5c3_Err
	})
}
