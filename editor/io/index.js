import fetch from 'isomorphic-unfetch'

const dummyState = {
    projectId: '5b1c2e345d9b1d61551da093',
}

export const fetchPageHierarchy = () => {
    return fetch(`${process.env.BASE_URL}/api/projects/${dummyState.projectId}/pages/`)
}

export const fetchTemplates = () => {
    return fetch(`${process.env.BASE_URL}/api/projects/${dummyState.projectId}/templates/`)
}

export const createTemplate = (data = {}) => {
    return fetch(`${process.env.BASE_URL}/api/projects/${dummyState.projectId}/templates/`, {
        body: JSON.stringify(data),
        headers: { 'content-type': 'application/json' },
        method: 'POST',
        mode: 'cors'
    })
}