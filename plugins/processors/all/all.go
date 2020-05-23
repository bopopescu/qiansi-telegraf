package all

import (
	_ "gitee.com/zhimiao/qiansi-telegraf/plugins/processors/clone"
	_ "gitee.com/zhimiao/qiansi-telegraf/plugins/processors/converter"
	_ "gitee.com/zhimiao/qiansi-telegraf/plugins/processors/date"
	_ "gitee.com/zhimiao/qiansi-telegraf/plugins/processors/dedup"
	_ "gitee.com/zhimiao/qiansi-telegraf/plugins/processors/defaults"
	_ "gitee.com/zhimiao/qiansi-telegraf/plugins/processors/enum"
	_ "gitee.com/zhimiao/qiansi-telegraf/plugins/processors/filepath"
	_ "gitee.com/zhimiao/qiansi-telegraf/plugins/processors/override"
	_ "gitee.com/zhimiao/qiansi-telegraf/plugins/processors/parser"
	_ "gitee.com/zhimiao/qiansi-telegraf/plugins/processors/pivot"
	_ "gitee.com/zhimiao/qiansi-telegraf/plugins/processors/printer"
	_ "gitee.com/zhimiao/qiansi-telegraf/plugins/processors/regex"
	_ "gitee.com/zhimiao/qiansi-telegraf/plugins/processors/rename"
	_ "gitee.com/zhimiao/qiansi-telegraf/plugins/processors/s2geo"
	_ "gitee.com/zhimiao/qiansi-telegraf/plugins/processors/strings"
	_ "gitee.com/zhimiao/qiansi-telegraf/plugins/processors/tag_limit"
	_ "gitee.com/zhimiao/qiansi-telegraf/plugins/processors/template"
	_ "gitee.com/zhimiao/qiansi-telegraf/plugins/processors/topk"
	_ "gitee.com/zhimiao/qiansi-telegraf/plugins/processors/unpivot"
)
