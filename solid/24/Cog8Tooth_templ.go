// Code generated by templ - DO NOT EDIT.

// templ: version: 0.2.476
package solid24

//lint:file-ignore SA4006 This context is only used if a nested component is present.

import "github.com/a-h/templ"
import "context"
import "io"
import "bytes"

func Cog8ToothIcon() templ.Component {
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
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<svg width=\"24\" height=\"24\" viewBox=\"0 0 24 24\" fill=\"currentColor\" xmlns=\"http://www.w3.org/2000/svg\" class=\"w-6 h-6\"><path fill-rule=\"evenodd\" clip-rule=\"evenodd\" d=\"M11.828 2.25C10.9115 2.25 10.1292 2.91265 9.97854 3.81675L9.88699 4.36602C9.84017 4.64695 9.63707 4.87407 9.37035 4.97395C9.20811 5.0347 9.04853 5.1009 8.89184 5.17235C8.63259 5.29056 8.32822 5.27363 8.09636 5.10801L7.64299 4.78418C6.89714 4.25143 5.87546 4.33599 5.22734 4.9841L4.98413 5.22732C4.33601 5.87544 4.25145 6.89712 4.7842 7.64297L5.10803 8.09633C5.27364 8.32819 5.29057 8.63256 5.17236 8.89182C5.10091 9.04851 5.0347 9.20809 4.97395 9.37034C4.87408 9.63706 4.64695 9.84016 4.36602 9.88698L3.81675 9.97852C2.91265 10.1292 2.25 10.9114 2.25 11.828V12.172C2.25 13.0885 2.91265 13.8708 3.81675 14.0215L4.36602 14.113C4.64695 14.1598 4.87407 14.3629 4.97395 14.6297C5.03469 14.7919 5.1009 14.9515 5.17234 15.1082C5.29056 15.3674 5.27362 15.6718 5.10801 15.9036L4.78416 16.357C4.25141 17.1029 4.33597 18.1246 4.98408 18.7727L5.2273 19.0159C5.87541 19.664 6.8971 19.7486 7.64295 19.2158L8.09633 18.892C8.32819 18.7264 8.63256 18.7094 8.89182 18.8276C9.04851 18.8991 9.20809 18.9653 9.37034 19.026C9.63706 19.1259 9.84016 19.353 9.88698 19.634L9.97852 20.1832C10.1292 21.0873 10.9114 21.75 11.828 21.75H12.172C13.0885 21.75 13.8708 21.0874 14.0215 20.1832L14.113 19.634C14.1598 19.3531 14.3629 19.1259 14.6297 19.0261C14.7919 18.9653 14.9515 18.8991 15.1082 18.8277C15.3674 18.7094 15.6718 18.7264 15.9036 18.892L16.357 19.2158C17.1029 19.7486 18.1245 19.664 18.7727 19.0159L19.0159 18.7727C19.664 18.1246 19.7485 17.1029 19.2158 16.357L18.892 15.9037C18.7264 15.6718 18.7094 15.3674 18.8276 15.1082C18.8991 14.9515 18.9653 14.7919 19.026 14.6297C19.1259 14.3629 19.353 14.1598 19.634 14.113L20.1832 14.0215C21.0873 13.8708 21.75 13.0886 21.75 12.172V11.828C21.75 10.9115 21.0874 10.1292 20.1832 9.97854L19.634 9.88699C19.3531 9.84017 19.1259 9.63707 19.0261 9.37035C18.9653 9.20811 18.8991 9.04854 18.8277 8.89185C18.7094 8.63259 18.7264 8.32822 18.892 8.09636L19.2158 7.64297C19.7486 6.89712 19.664 5.87544 19.0159 5.22732L18.7727 4.9841C18.1246 4.33599 17.1029 4.25143 16.3571 4.78418L15.9037 5.10802C15.6718 5.27364 15.3674 5.29057 15.1082 5.17236C14.9515 5.10091 14.7919 5.0347 14.6297 4.97395C14.3629 4.87408 14.1598 4.64695 14.113 4.36602L14.0215 3.81675C13.8708 2.91265 13.0886 2.25 12.172 2.25H11.828ZM12 15.75C14.0711 15.75 15.75 14.0711 15.75 12C15.75 9.92893 14.0711 8.25 12 8.25C9.92893 8.25 8.25 9.92893 8.25 12C8.25 14.0711 9.92893 15.75 12 15.75Z\" fill=\"#0F172A\"></path></svg>")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		if !templ_7745c5c3_IsBuffer {
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteTo(templ_7745c5c3_W)
		}
		return templ_7745c5c3_Err
	})
}
