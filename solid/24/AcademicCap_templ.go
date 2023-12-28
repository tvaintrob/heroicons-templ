// Code generated by templ - DO NOT EDIT.

// templ: version: 0.2.476
package solid24

//lint:file-ignore SA4006 This context is only used if a nested component is present.

import "github.com/a-h/templ"
import "context"
import "io"
import "bytes"

func AcademicCapIcon() templ.Component {
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
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<svg width=\"24\" height=\"24\" viewBox=\"0 0 24 24\" fill=\"none\" xmlns=\"http://www.w3.org/2000/svg\"><path d=\"M11.6998 2.80529C11.8912 2.72164 12.1089 2.72164 12.3003 2.80529C16.0192 4.43011 19.5437 6.41637 22.8295 8.71956C23.0673 8.88623 23.1875 9.1752 23.1381 9.46135C23.0887 9.7475 22.8785 9.97941 22.5986 10.0567C21.9137 10.2457 21.2347 10.4494 20.5618 10.6663C17.8307 11.5471 15.2018 12.6554 12.6972 13.9688L12.6939 13.9705C12.5803 14.0301 12.467 14.09 12.354 14.1504C12.1331 14.2684 11.8679 14.2684 11.6471 14.1504C11.533 14.0895 11.4186 14.0289 11.3039 13.9688C10.0655 13.3193 8.79658 12.7201 7.5 12.1736V11.95C7.5 11.8186 7.56742 11.702 7.67173 11.6389C9.17685 10.727 10.7294 9.88565 12.3247 9.11936C12.6981 8.94002 12.8554 8.49195 12.6761 8.11858C12.4967 7.7452 12.0486 7.58791 11.6753 7.76725C10.036 8.55463 8.44086 9.41909 6.89449 10.3559C6.44111 10.6306 6.13632 11.0801 6.03607 11.5838C5.18115 11.2549 4.31499 10.9486 3.43829 10.6659C2.76546 10.4489 2.08644 10.2457 1.40154 10.0567C1.12162 9.9794 0.911461 9.74749 0.86204 9.46134C0.812619 9.17519 0.932824 8.88622 1.17061 8.71955C4.45645 6.41636 7.98097 4.43011 11.6998 2.80529Z\" fill=\"#0F172A\"></path> <path d=\"M13.0609 15.4734C15.4997 14.1703 18.0621 13.0687 20.7258 12.1906C20.8601 13.6054 20.9458 15.0343 20.9813 16.4755C20.9889 16.7847 20.8059 17.0669 20.5205 17.1861C17.6693 18.3764 14.9574 19.834 12.4159 21.5277C12.1641 21.6955 11.836 21.6955 11.5841 21.5277C9.04267 19.834 6.33073 18.3764 3.4796 17.1861C3.19416 17.0669 3.01116 16.7847 3.01878 16.4755C3.05429 15.0342 3.14001 13.6052 3.27427 12.1903C4.19527 12.4938 5.10415 12.8242 6 13.1803V14.4507C5.55165 14.71 5.25 15.1948 5.25 15.75C5.25 16.2453 5.49008 16.6846 5.86022 16.9577C5.7707 17.3383 5.63822 17.7108 5.46277 18.0675C5.91546 18.2811 6.36428 18.5017 6.8091 18.7289C7.06243 18.2137 7.24612 17.6729 7.36014 17.1207C7.88449 16.887 8.25 16.3612 8.25 15.75C8.25 15.1948 7.94835 14.71 7.5 14.4507V13.8059C8.6714 14.3177 9.81885 14.8743 10.9402 15.4734C11.6028 15.8274 12.3983 15.8274 13.0609 15.4734Z\" fill=\"#0F172A\"></path> <path d=\"M4.46222 19.4623C4.88136 19.0432 5.21502 18.5711 5.46277 18.0675C5.91546 18.2811 6.36428 18.5017 6.8091 18.7289C6.49055 19.3768 6.06164 19.9842 5.52288 20.523C5.22999 20.8158 4.75512 20.8158 4.46222 20.523C4.16933 20.2301 4.16933 19.7552 4.46222 19.4623Z\" fill=\"#0F172A\"></path></svg>")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		if !templ_7745c5c3_IsBuffer {
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteTo(templ_7745c5c3_W)
		}
		return templ_7745c5c3_Err
	})
}
