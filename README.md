Simple go example not using the standard golang community folder convention.
```
[application]<---[infrastructure]--->[domain]
|                                         ^
|_________________________________________|
```
- domain is protected

- application used by delivery mechanism (infrastructure)

- infrastructure isolated from polluting application or domain

````go
func (cm *CreateMenu) Create(createMenuRequest CreateMenuRequest, createMenuPresenter CreateMenuPresenter) {

    menu := domain.NewMenu(
        domain.NewCafeId(domain.NewUuid()),
        domain.NewTitle(createMenuRequest.Title),
    )

    createMenuPresenter.Present(
        CreateMenuResponse{
            Uuid:  menu.GetCafeId().GetValue(),
            Title: menu.GetTitle().GetValue(),
        },
    )
}
````

