// Code generated by templ - DO NOT EDIT.

// templ: version: 0.2.476
package solid24

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
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<svg width=\"24\" height=\"24\" viewBox=\"0 0 24 24\" fill=\"currentColor\" xmlns=\"http://www.w3.org/2000/svg\" class=\"w-6 h-6\"><path d=\"M15 1.784L14.2045 2.5795C13.7652 3.01884 13.7652 3.73115 14.2045 4.17049C14.6438 4.60983 15.3562 4.60983 15.7955 4.17049C16.2348 3.73115 16.2348 3.01884 15.7955 2.5795L15 1.784Z\" fill=\"#0F172A\"></path> <path d=\"M12 1.784L11.2045 2.5795C10.7652 3.01884 10.7652 3.73115 11.2045 4.17049C11.6438 4.60983 12.3562 4.60983 12.7955 4.17049C13.2348 3.73115 13.2348 3.01884 12.7955 2.5795L12 1.784Z\" fill=\"#0F172A\"></path> <path d=\"M8.99999 1.784L8.2045 2.5795C7.76516 3.01884 7.76516 3.73115 8.2045 4.17049C8.64384 4.60983 9.35616 4.60983 9.7955 4.17049C10.2348 3.73115 10.2348 3.01884 9.7955 2.5795L8.99999 1.784Z\" fill=\"#0F172A\"></path> <path d=\"M9.75 7.54669C10.2483 7.52597 10.7483 7.5121 11.25 7.50518V6.75C11.25 6.33579 11.5858 6 12 6C12.4142 6 12.75 6.33579 12.75 6.75V7.50518C13.2517 7.5121 13.7517 7.52597 14.25 7.54669V6.75C14.25 6.33579 14.5858 6 15 6C15.4142 6 15.75 6.33579 15.75 6.75V7.63003C15.8524 7.63715 15.9547 7.64456 16.0569 7.65226C17.6071 7.76907 18.75 9.07932 18.75 10.5976V11.6162C16.5333 11.3742 14.2811 11.25 12 11.25C9.71886 11.25 7.46673 11.3742 5.25 11.6162V10.5976C5.25 9.07932 6.39295 7.76907 7.94314 7.65226C8.04534 7.64456 8.14763 7.63715 8.25 7.63003V6.75C8.25 6.33579 8.58579 6 9 6C9.41421 6 9.75 6.33579 9.75 6.75V7.54669Z\" fill=\"#0F172A\"></path> <path d=\"M12 12.75C9.52847 12.75 7.09944 12.934 4.72596 13.2891C3.27191 13.5067 2.25 14.7716 2.25 16.2057V16.59C3.11853 16.4286 4.02704 16.55 4.83541 16.9542C5.56854 17.3207 6.43146 17.3207 7.16459 16.9542C8.32001 16.3765 9.67999 16.3765 10.8354 16.9542C11.5685 17.3207 12.4315 17.3207 13.1646 16.9542C14.32 16.3765 15.68 16.3765 16.8354 16.9542C17.5685 17.3207 18.4315 17.3207 19.1646 16.9542C19.973 16.55 20.8815 16.4286 21.75 16.59V16.2057C21.75 14.7716 20.7281 13.5067 19.274 13.2891C16.9006 12.934 14.4715 12.75 12 12.75Z\" fill=\"#0F172A\"></path> <path d=\"M21.75 18.1312C21.1195 17.9416 20.4342 17.9964 19.8354 18.2958C18.68 18.8735 17.32 18.8735 16.1646 18.2958C15.4315 17.9293 14.5685 17.9293 13.8354 18.2958C12.68 18.8735 11.32 18.8735 10.1646 18.2958C9.43146 17.9293 8.56854 17.9293 7.83541 18.2958C6.67999 18.8735 5.32001 18.8735 4.16459 18.2958C3.56583 17.9964 2.88049 17.9416 2.25 18.1312V20.625C2.25 21.6605 3.08947 22.5 4.125 22.5H19.875C20.9105 22.5 21.75 21.6605 21.75 20.625V18.1312Z\" fill=\"#0F172A\"></path></svg>")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		if !templ_7745c5c3_IsBuffer {
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteTo(templ_7745c5c3_W)
		}
		return templ_7745c5c3_Err
	})
}
