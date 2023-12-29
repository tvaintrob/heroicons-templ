// Code generated by templ - DO NOT EDIT.

// templ: version: 0.2.476
package solid16

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
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<svg width=\"16\" height=\"16\" viewBox=\"0 0 16 16\" fill=\"currentColor\" xmlns=\"http://www.w3.org/2000/svg\" class=\"w-4 h-4\"><path d=\"M7.70177 1.36758C7.89214 1.28502 8.10821 1.28502 8.29858 1.36758C10.397 2.27761 12.4042 3.35793 14.3029 4.59105C14.548 4.75028 14.6783 5.03746 14.6367 5.32682C14.5951 5.61618 14.3891 5.85498 14.109 5.93866C12.1038 6.53768 10.1751 7.31422 8.34107 8.25008C8.12694 8.35934 7.87341 8.35934 7.65928 8.25008C7.03364 7.93083 6.39698 7.63012 5.75002 7.34867V6.80657C6.58634 6.33143 7.44297 5.88795 8.31827 5.47777C8.69334 5.302 8.85491 4.85546 8.67914 4.48039C8.50338 4.10531 8.05683 3.94375 7.68176 4.11951C6.63525 4.60993 5.61447 5.14603 4.62198 5.72525C4.39164 5.85967 4.25002 6.10631 4.25002 6.37301V6.73785C3.47668 6.44405 2.69006 6.17727 1.8913 5.93866C1.6112 5.85498 1.40524 5.61618 1.36362 5.32682C1.322 5.03746 1.4523 4.75028 1.69747 4.59105C3.59612 3.35793 5.60336 2.27761 7.70177 1.36758Z\" fill=\"#0F172A\"></path> <path d=\"M4.25002 8.34775C3.71965 8.1356 3.18252 7.93682 2.63904 7.75185C2.46449 8.73026 2.32464 9.72067 2.22082 10.7218C2.18598 11.0578 2.38004 11.3756 2.69483 11.4981C2.87032 11.5664 3.04499 11.6363 3.21883 11.7078C3.04706 11.9465 2.85382 12.1744 2.6391 12.3892C2.34621 12.682 2.34621 13.1569 2.6391 13.4498C2.932 13.7427 3.40687 13.7427 3.69976 13.4498C4.04924 13.1003 4.35406 12.723 4.61419 12.325C4.15544 12.1076 3.6902 11.9018 3.21883 11.7078C3.90636 10.7527 4.25001 9.62721 4.25002 8.5001V8.34775Z\" fill=\"#0F172A\"></path> <path d=\"M7.60299 13.9591C6.64228 13.3592 5.64461 12.8131 4.61419 12.325C5.27965 11.3067 5.65275 10.1531 5.73335 8.98392C6.15287 9.17619 6.56767 9.37701 6.97754 9.58616C7.61992 9.91395 8.38051 9.91395 9.02289 9.58616C10.4151 8.87576 11.8641 8.26144 13.3614 7.75183C13.536 8.73024 13.6758 9.72066 13.7797 10.7218C13.8145 11.0578 13.6204 11.3756 13.3057 11.4981C11.5856 12.1674 9.94367 12.9936 8.39748 13.9591C8.1544 14.1109 7.84608 14.1109 7.60299 13.9591Z\" fill=\"#0F172A\"></path></svg>")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		if !templ_7745c5c3_IsBuffer {
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteTo(templ_7745c5c3_W)
		}
		return templ_7745c5c3_Err
	})
}
