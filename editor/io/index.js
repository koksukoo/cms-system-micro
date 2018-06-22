import fetch from 'isomorphic-unfetch'

const dummyState = {
    projectId: '5b1c2e345d9b1d61551da093',
}

export const fetchPageHierarchy = () => {
    return fetch(`http://localhost:3000/fetch/pages/`)
}