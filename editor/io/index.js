import fetch from 'isomorphic-unfetch'

const dummyState = {
    projectId: '5b6ec7c45d9b1d4133a2b24d',
}

export const fetchPageHierarchy = () => {
    return fetch(`${process.env.BASE_URL}/api/projects/${dummyState.projectId}/pages/`)
}

export const fetchTemplates = () => {
    return fetch(`${process.env.BASE_URL}/api/projects/${dummyState.projectId}/templates/`)
}

export const fetchTemplate = (id) => {
    return fetch(`${process.env.BASE_URL}/api/projects/${dummyState.projectId}/templates/${id}`)
}

export const createTemplate = (data = {}) => {
    return fetch(`${process.env.BASE_URL}/api/projects/${dummyState.projectId}/templates/`, {
        body: JSON.stringify(data),
        headers: { 'content-type': 'application/json' },
        method: 'POST',
        mode: 'cors'
    })
}

export const updateTemplate = (id, data = {}) => {
    return fetch(`${process.env.BASE_URL}/api/projects/${dummyState.projectId}/templates/${id}`, {
        body: JSON.stringify(data),
        headers: { 'content-type': 'application/json' },
        method: 'PUT',
        mode: 'cors'
    })
}

export const deleteTemplate = (id) => {
    return fetch(`${process.env.BASE_URL}/api/projects/${dummyState.projectId}/templates/${id}`, {
        method: 'DELETE',
    })
}