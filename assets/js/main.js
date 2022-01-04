console.log("main.js loaded")
// sidebar
const sidebarShowClassName = "show-sidebar"
const sidebarCloseClassName = "hide-sidebar"
const sidebar = document.getElementById("sideNavBar")
const openSideBarTogler = document.getElementById("sideBarTogler")
const closeSideBarTogler = document.getElementById("sideBarToglerClose")
closeSideBarTogler.onclick = () => {
    sidebar.className = sidebar.className.replace(sidebarShowClassName, sidebarCloseClassName)
}
openSideBarTogler.onclick = () => {
    sidebar.className = sidebar.className.replace(sidebarCloseClassName, sidebarShowClassName)
}
// end
