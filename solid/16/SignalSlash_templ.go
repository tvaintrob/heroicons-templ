// Code generated by templ - DO NOT EDIT.

// templ: version: 0.2.476
package solid16

//lint:file-ignore SA4006 This context is only used if a nested component is present.

import "github.com/a-h/templ"
import "context"
import "io"
import "bytes"

func SignalSlashIcon() templ.Component {
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
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<svg width=\"16\" height=\"16\" viewBox=\"0 0 16 16\" fill=\"none\" xmlns=\"http://www.w3.org/2000/svg\"><path d=\"M3.28033 2.21967C2.98744 1.92678 2.51256 1.92678 2.21967 2.21967C1.92678 2.51256 1.92678 2.98744 2.21967 3.28033L7.00193 8.06259C7.03295 8.56502 7.43498 8.96705 7.93741 8.99807L12.7197 13.7803C13.0126 14.0732 13.4874 14.0732 13.7803 13.7803C14.0732 13.4874 14.0732 13.0126 13.7803 12.7197L8.99807 7.93741C8.96705 7.43498 8.56502 7.03295 8.06259 7.00193L3.28033 2.21967Z\" fill=\"black\"></path> <path d=\"M3.05025 12.9502C0.873058 10.773 0.42986 7.51846 1.72066 4.90264L2.85932 6.0413C2.11257 8.00408 2.52977 10.3084 4.11091 11.8896C4.40381 12.1825 4.40381 12.6573 4.11091 12.9502C3.81802 13.2431 3.34315 13.2431 3.05025 12.9502Z\" fill=\"black\"></path> <path d=\"M5.25996 10.74C4.33993 9.82 3.97925 8.55243 4.17793 7.35991L5.86994 9.05192C5.98183 9.27844 6.13206 9.49082 6.32062 9.67938C6.61351 9.97227 6.61351 10.4471 6.32062 10.74C6.02773 11.0329 5.55285 11.0329 5.25996 10.74Z\" fill=\"black\"></path> <path d=\"M12.9497 3.05072C15.1268 5.22774 15.5701 8.48196 14.2797 11.0977L13.1409 9.95896C13.8873 7.99633 13.4701 5.69235 11.8891 4.11138C11.5962 3.81849 11.5962 3.34362 11.8891 3.05072C12.182 2.75783 12.6569 2.75783 12.9497 3.05072Z\" fill=\"black\"></path> <path d=\"M10.74 5.25996C11.6601 6.18 12.0208 7.44757 11.8221 8.64009L10.1301 6.94808C10.0182 6.72156 9.86794 6.50918 9.67938 6.32062C9.38649 6.02773 9.38649 5.55285 9.67938 5.25996C9.97227 4.96707 10.4471 4.96707 10.74 5.25996Z\" fill=\"black\"></path></svg>")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		if !templ_7745c5c3_IsBuffer {
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteTo(templ_7745c5c3_W)
		}
		return templ_7745c5c3_Err
	})
}
