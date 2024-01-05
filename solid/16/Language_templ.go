// Code generated by templ - DO NOT EDIT.

// templ: version: v0.2.513
package solid16

//lint:file-ignore SA4006 This context is only used if a nested component is present.

import "github.com/a-h/templ"
import "context"
import "io"
import "bytes"

func LanguageIcon(attrs templ.Attributes) templ.Component {
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
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<svg xmlns=\"http://www.w3.org/2000/svg\" fill=\"currentColor\" viewBox=\"0 0 16 16\" class=\"w-4 h-4\"")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		templ_7745c5c3_Err = templ.RenderAttributes(ctx, templ_7745c5c3_Buffer, attrs)
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("><path fill-rule=\"evenodd\" clip-rule=\"evenodd\" d=\"M11 5C11.299 5 11.5693 5.17751 11.6882 5.45179L14.9382 12.9518C15.1029 13.3319 14.9283 13.7735 14.5482 13.9382C14.1682 14.1029 13.7266 13.9283 13.5619 13.5482L12.8908 11.9997H9.10923L8.4382 13.5482C8.2735 13.9283 7.83189 14.1029 7.45182 13.9382C7.07176 13.7735 6.89717 13.3319 7.06186 12.9518L10.3119 5.45179C10.4307 5.17751 10.7011 5 11 5ZM9.75923 10.4997H12.2408L11 7.63628L9.75923 10.4997Z\"></path><path fill-rule=\"evenodd\" clip-rule=\"evenodd\" d=\"M5.00003 1C5.41424 1 5.75003 1.33579 5.75003 1.75V3.01104C6.16299 3.02322 6.5735 3.04541 6.98131 3.0774C7.44038 3.11341 7.89601 3.16182 8.34786 3.22231C8.75842 3.27727 9.04668 3.65464 8.99172 4.06519C8.93676 4.47574 8.55938 4.76401 8.14883 4.70905C7.92894 4.67961 7.70808 4.65321 7.48628 4.6299C7.1301 5.85717 6.59808 7.00928 5.91941 8.05729C6.15555 8.36066 6.40658 8.65193 6.67142 8.92999C6.95709 9.22993 6.94553 9.70466 6.64559 9.99034C6.34565 10.276 5.87092 10.2644 5.58525 9.96451C5.38294 9.7521 5.18774 9.53284 5.00002 9.30711C4.18402 10.2884 3.22645 11.1474 2.15883 11.853C1.81326 12.0813 1.34799 11.9863 1.11962 11.6408C0.891239 11.2952 0.986242 10.8299 1.33181 10.6015C2.3813 9.90797 3.31021 9.04714 4.08066 8.05729C3.88359 7.75296 3.69887 7.43984 3.52724 7.11865C3.33202 6.75332 3.46992 6.29891 3.83524 6.10369C4.20057 5.90847 4.65498 6.04637 4.8502 6.4117C4.89895 6.50293 4.9489 6.59343 5.00002 6.68318C5.38798 6.00207 5.7083 5.27759 5.95187 4.51891C5.63619 4.50635 5.31887 4.5 5.00003 4.5C3.93193 4.5 2.88086 4.57121 1.85122 4.70905C1.44067 4.76401 1.0633 4.47574 1.00834 4.06519C0.95338 3.65464 1.24164 3.27727 1.65219 3.22231C2.50548 3.10808 3.37219 3.03692 4.25003 3.01104V1.75C4.25003 1.33579 4.58582 1 5.00003 1Z\"></path></svg>")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		if !templ_7745c5c3_IsBuffer {
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteTo(templ_7745c5c3_W)
		}
		return templ_7745c5c3_Err
	})
}
