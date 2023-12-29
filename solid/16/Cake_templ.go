// Code generated by templ - DO NOT EDIT.

// templ: version: 0.2.476
package solid16

//lint:file-ignore SA4006 This context is only used if a nested component is present.

import "github.com/a-h/templ"
import "context"
import "io"
import "bytes"

func CakeIcon() templ.Component {
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
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<svg width=\"16\" height=\"16\" viewBox=\"0 0 16 16\" fill=\"currentColor\" xmlns=\"http://www.w3.org/2000/svg\" class=\"w-4 h-4\"><path d=\"M4.74999 1L3.86612 1.88389C3.37796 2.37205 3.37796 3.16351 3.86612 3.65166C4.35427 4.13982 5.14573 4.13982 5.63388 3.65166C6.12204 3.16351 6.12204 2.37205 5.63388 1.88389L4.74999 1Z\"></path> <path d=\"M11.25 1L10.3661 1.88389C9.87796 2.37205 9.87796 3.16351 10.3661 3.65166C10.8543 4.13982 11.6457 4.13982 12.1339 3.65166C12.622 3.16351 12.622 2.37205 12.1339 1.88389L11.25 1Z\"></path> <path d=\"M8.88388 1.88389L7.99999 1L7.11612 1.88389C6.62796 2.37205 6.62796 3.16351 7.11612 3.65166C7.60427 4.13982 8.39573 4.13982 8.88388 3.65166C9.37204 3.16351 9.37204 2.37205 8.88388 1.88389Z\"></path> <path d=\"M4 7C2.89543 7 2 7.89543 2 9V10.0338C2.34716 10.0339 2.69432 9.97826 3.02794 9.86705L3.49812 9.71033C4.47314 9.38532 5.52727 9.38532 6.50229 9.71033L6.97246 9.86705C7.63958 10.0894 8.36083 10.0894 9.02794 9.86705L9.49812 9.71033C10.4731 9.38532 11.5273 9.38532 12.5023 9.71033L12.9725 9.86705C13.306 9.97822 13.653 10.0338 14 10.0338V9C14 7.89543 13.1046 7 12 7V5.75C12 5.33579 11.6642 5 11.25 5C10.8358 5 10.5 5.33579 10.5 5.75V7H8.75V5.75C8.75 5.33579 8.41421 5 8 5C7.58579 5 7.25 5.33579 7.25 5.75V7H5.5V5.75C5.5 5.33579 5.16421 5 4.75 5C4.33579 5 4 5.33579 4 5.75V7Z\"></path> <path d=\"M14 11.5338C13.4928 11.5338 12.9856 11.4526 12.4981 11.2901L12.0279 11.1334C11.3608 10.911 10.6396 10.911 9.97246 11.1334L9.50229 11.2901C8.52727 11.6151 7.47314 11.6151 6.49812 11.2901L6.02794 11.1334C5.36083 10.911 4.63958 10.911 3.97246 11.1334L3.50229 11.2901C3.01471 11.4526 2.50735 11.5339 2 11.5338V13C2 13.5523 2.44772 14 3 14H13C13.5523 14 14 13.5523 14 13V11.5338Z\"></path></svg>")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		if !templ_7745c5c3_IsBuffer {
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteTo(templ_7745c5c3_W)
		}
		return templ_7745c5c3_Err
	})
}
