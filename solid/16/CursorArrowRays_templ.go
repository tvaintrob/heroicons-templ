// Code generated by templ - DO NOT EDIT.

// templ: version: 0.2.476
package solid16

//lint:file-ignore SA4006 This context is only used if a nested component is present.

import "github.com/a-h/templ"
import "context"
import "io"
import "bytes"

func CursorArrowRaysIcon() templ.Component {
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
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<svg width=\"16\" height=\"16\" viewBox=\"0 0 16 16\" fill=\"none\" xmlns=\"http://www.w3.org/2000/svg\" class=\"w-4 h-4\"><path d=\"M7.25 1.75C7.25 1.33579 7.58579 1 8 1C8.41421 1 8.75 1.33579 8.75 1.75V3.25C8.75 3.66421 8.41421 4 8 4C7.58579 4 7.25 3.66421 7.25 3.25V1.75Z\" fill=\"black\"></path> <path d=\"M11.5356 2.9038C11.8285 2.61091 12.3034 2.61091 12.5962 2.9038C12.8891 3.1967 12.8891 3.67157 12.5962 3.96446L11.5356 5.02512C11.2427 5.31802 10.7678 5.31802 10.4749 5.02512C10.182 4.73223 10.182 4.25736 10.4749 3.96446L11.5356 2.9038Z\" fill=\"black\"></path> <path d=\"M14.5 7.5C14.5 7.08579 14.1642 6.75 13.75 6.75H12.25C11.8358 6.75 11.5 7.08579 11.5 7.5C11.5 7.91421 11.8358 8.25 12.25 8.25H13.75C14.1642 8.25 14.5 7.91421 14.5 7.5Z\" fill=\"black\"></path> <path d=\"M4.46442 9.97488C4.75731 9.68199 5.23219 9.68199 5.52508 9.97488C5.81797 10.2678 5.81797 10.7426 5.52508 11.0355L4.46442 12.0962C4.17153 12.3891 3.69665 12.3891 3.40376 12.0962C3.11087 11.8033 3.11087 11.3284 3.40376 11.0355L4.46442 9.97488Z\" fill=\"black\"></path> <path d=\"M4.5 7.5C4.5 7.08579 4.16421 6.75 3.75 6.75H2.25C1.83579 6.75 1.5 7.08579 1.5 7.5C1.5 7.91421 1.83579 8.25 2.25 8.25H3.75C4.16421 8.25 4.5 7.91421 4.5 7.5Z\" fill=\"black\"></path> <path d=\"M5.52509 3.96447C5.81798 4.25736 5.81798 4.73223 5.52509 5.02513C5.2322 5.31802 4.75732 5.31802 4.46443 5.02513L3.40377 3.96447C3.11088 3.67157 3.11088 3.1967 3.40377 2.90381C3.69666 2.61091 4.17154 2.61091 4.46443 2.90381L5.52509 3.96447Z\" fill=\"black\"></path> <path d=\"M8.77875 7.438C8.59919 7.17468 8.27284 7.05361 7.96499 7.13609C7.65714 7.21858 7.43504 7.48662 7.4112 7.80443L7.01493 13.087C6.99268 13.3835 7.14774 13.6652 7.41019 13.805C7.67265 13.9448 7.99292 13.9164 8.22662 13.7325L8.82888 13.2586L9.11664 14.3332C9.22379 14.7333 9.63501 14.9708 10.0351 14.8636C10.4352 14.7565 10.6727 14.3453 10.5656 13.9451L10.2777 12.8702L11.0366 12.9796C11.3309 13.022 11.6225 12.8865 11.7799 12.6342C11.9373 12.3819 11.9307 12.0604 11.7632 11.8147L8.77875 7.438Z\" fill=\"black\"></path></svg>")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		if !templ_7745c5c3_IsBuffer {
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteTo(templ_7745c5c3_W)
		}
		return templ_7745c5c3_Err
	})
}
