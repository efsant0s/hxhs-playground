package handler

import (
	"hxhs-playground/component"
	"hxhs-playground/view"

	"github.com/a-h/templ"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/adaptor"
	"github.com/google/uuid"
)

func HandleTestPage(c *fiber.Ctx) error {
	page := view.DemoPage()
	handler := adaptor.HTTPHandler(templ.Handler(page))
	return handler(c)
}
func DemoPageView(c *fiber.Ctx) error {
	pageContent := view.DemoPageContent{
		PageTitle: "Component Demo Page",
		PageID:    "demo-page-container",
		Groups: []view.ComponentGroup{
			getActionBarGroup(),
			getBPMNGroup(),
			getButtonGroup(),
			getCardGroup(),
			getContainerGroup(),
			getDatatableGroup(),
			getFormGroup(),
			getInputGroup(),
			getModalGroup(),
			getSelectGroup(),
			getSwalFireGroup()},
	}
	page := view.MainDemoPage(pageContent)
	handler := adaptor.HTTPHandler(templ.Handler(page))
	return handler(c)
}

func getActionBarGroup() view.ComponentGroup {
	return view.ComponentGroup{
		GroupName:  "Action Bar Components",
		GroupID:    "action-bar-group",
		Components: []templ.Component{},
	}
}

func getBPMNGroup() view.ComponentGroup {
	return view.ComponentGroup{
		GroupName:  "BPMN Components",
		GroupID:    "bpmn-group",
		Components: []templ.Component{},
	}
}

func getButtonGroup() view.ComponentGroup {
	defaultButton := component.ButtonComponent(component.DefaultButtonComponentItem(uuid.New().String(), "default-button", "Default Button"))
	buttonComponentType := component.ButtonComponent(component.ButtonComponentType(uuid.New().String(), "button-component-type", "Button with Type", "button"))
	buttonComponentClass := component.ButtonComponent(component.ButtonComponentClass(uuid.New().String(), "button-component-disabled", "Button with Class", "btn btn-primary"))
	ButtonComponentStyle := component.ButtonComponent(component.ButtonComponentStyle(uuid.New().String(), "button-component-disabled", "Button with Style", "background-color: #04AA6D;"))
	buttonComponentDisabled := component.ButtonComponent(component.ButtonComponentDisabled(uuid.New().String(), "button-component-disabled", "Button with Disabled", true))
	//
	buttonComponentHXGET := component.ButtonComponent(
		component.ButtonComponentHXGet(uuid.New().String(), "button-component-hx-get", "Button with HX-GET", "api/getButtonComponentHX", ""))
	buttonComponentHXGETHover := component.ButtonComponent(
		component.ButtonComponentHXGet(uuid.New().String(), "button-component-hx-get-on-hover", "Button with HX-GET ON dbclick, not click", "api/getButtonComponentHX", "dblclick"))
	buttonComponentHXPOST := component.ButtonComponent(
		component.ButtonComponentHXPost(uuid.New().String(), "button-component-hx-post", "Button with HX-POST", "api/postButtonComponentHX", ""))
	hxDIVId := "id" + "DIVTESTHX"
	buttonComponentHXGetTarget := component.ButtonComponent(
		component.ButtonComponentHXGetTarget(uuid.New().String(), "button-component-hx-get-target", "Button with HX-GET Target", "api/getButtonComponentHXTarget", hxDIVId))

	divWithButton := component.GenerateDivWitComponent(hxDIVId, buttonComponentHXGetTarget)
	buttonComponentHXGetSwap := component.ButtonComponent(component.ButtonComponentHXGetSwap(uuid.New().String(), "button-component-hx-get-swap-afterend", "Button with HX-GET Swap afterend", "api/getButtonComponentHXSwapAfterend", "afterend"))

	idDivInclude := "id" + "DIVTESTHXINCLUDE"
	buttonComponentHXGetInclude := component.ButtonComponent(
		component.ButtonComponentHXPostInclude(uuid.New().String(), "button-component-hx-post-include", "Button with HX-POST Include", "api/postButtonComponentHXInclude", idDivInclude))
	divWithInputForButtonInclude := component.GenerateDivWithInputAndComponent(idDivInclude, buttonComponentHXGetInclude)

	return view.ComponentGroup{
		GroupName: "Button Components",
		GroupID:   "button-group",
		Components: []templ.Component{defaultButton, buttonComponentType, buttonComponentClass, ButtonComponentStyle,
			buttonComponentDisabled, buttonComponentHXGET, buttonComponentHXGETHover, buttonComponentHXPOST, divWithButton, buttonComponentHXGetSwap,
			divWithInputForButtonInclude},
	}
}

func getCardGroup() view.ComponentGroup {
	return view.ComponentGroup{
		GroupName:  "Card Components",
		GroupID:    "card-group",
		Components: []templ.Component{},
	}
}

func getContainerGroup() view.ComponentGroup {
	return view.ComponentGroup{
		GroupName:  "Container Components",
		GroupID:    "container-group",
		Components: []templ.Component{},
	}
}

func getDatatableGroup() view.ComponentGroup {
	return view.ComponentGroup{
		GroupName:  "Datatable Components",
		GroupID:    "datatable-group",
		Components: []templ.Component{},
	}
}

func getFormGroup() view.ComponentGroup {
	return view.ComponentGroup{
		GroupName:  "Form Components",
		GroupID:    "form-group",
		Components: []templ.Component{},
	}
}

func getInputGroup() view.ComponentGroup {
	return view.ComponentGroup{
		GroupName:  "Input Components",
		GroupID:    "input-group",
		Components: []templ.Component{},
	}
}

func getModalGroup() view.ComponentGroup {
	return view.ComponentGroup{
		GroupName:  "Modal Components",
		GroupID:    "modal-group",
		Components: []templ.Component{},
	}
}

func getSelectGroup() view.ComponentGroup {
	return view.ComponentGroup{
		GroupName:  "Select Components",
		GroupID:    "select-group",
		Components: []templ.Component{},
	}
}

func getSwalFireGroup() view.ComponentGroup {
	return view.ComponentGroup{
		GroupName:  "SwalFire Components",
		GroupID:    "swal-fire-group",
		Components: []templ.Component{},
	}
}
