import fetch from "isomorphic-unfetch"
import cookies from "js-cookie"

const dummyState = {
    projectId: "5b6ec7c45d9b1d4133a2b24d"
}

export const login = (data) => {
    return fetch(`${process.env.BASE_URL}/api/users/login`, {
        headers: { "content-type": "application/json" },
        method: "POST",
        body: JSON.stringify(data),
        mode: "cors",
    })
}

export const register = (data) => {
    return fetch(`${process.env.BASE_URL}/api/users/register`, {
        headers: { "content-type": "application/json" },
        method: "POST",
        body: JSON.stringify(data),
        mode: "cors",
    })
}

export const fetchPageHierarchy = () => {
    return fetch(`${process.env.BASE_URL}/api/projects/${dummyState.projectId}/pages/`)
}

export const fetchTemplates = () => {
    return fetch(`${process.env.BASE_URL}/api/projects/${dummyState.projectId}/templates/`)
}

export const fetchTemplate = id => {
    return fetch(`${process.env.BASE_URL}/api/projects/${dummyState.projectId}/templates/${id}`)
}

export const createTemplate = (data = {}) => {
    return fetch(`${process.env.BASE_URL}/api/projects/${dummyState.projectId}/templates/`, {
        body: JSON.stringify(data),
        headers: { "content-type": "application/json" },
        method: "POST",
        mode: "cors"
    })
}

export const updateTemplate = (id, data = {}) => {
    return fetch(`${process.env.BASE_URL}/api/projects/${dummyState.projectId}/templates/${id}`, {
        body: JSON.stringify(data),
        headers: { "content-type": "application/json" },
        method: "PUT",
        mode: "cors"
    })
}

export const deleteTemplate = id => {
    return fetch(`${process.env.BASE_URL}/api/projects/${dummyState.projectId}/templates/${id}`, {
        method: "DELETE"
    })
}

export const setCookie = (key, value) => {
    if (process.browser) {
        cookie.set(key, value, { expires: 1, path: "/" })
    }
}

export const getCookie = (key, req) => {
    return process.browser ? getCookieFromBrowser(key) : getCookieFromServer(key, req)
}

const getCookieFromBrowser = key => {
    return cookies.get(key)
}

const getCookieFromServer = (key, req) => {
    if (!Object.keys(req.cookies).length) {
      return undefined;
    }
    const value = req.cookies[key];
    return value;
  };
