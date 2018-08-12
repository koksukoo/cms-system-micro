import Router from "next/router"
import { getCookie } from "io"

export const redirect = (target, ctx = {}) => {
    if (ctx.res) {
        ctx.res.writeHead(302, { Location: target })
        ctx.res.end()
    } else {
        Router.replace(target)
    }
}

export const isAuthenticated = ctx => !!getCookie("cms_session", ctx.req)

export const redirectIfNotAuthenticated = ctx => {
    if (!isAuthenticated(ctx)) {
        redirect("/login", ctx)
        return true
    }
    return false
}