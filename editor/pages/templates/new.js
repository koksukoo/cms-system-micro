import { Component } from "react"
import dynamic from "next/dynamic"
import Layout from "components/BaseLayout"
import TemplateForm from "components/TemplateForm"
import "codemirror/lib/codemirror.css"

import { redirectIfNotAuthenticated } from "utils/helpers"

const Codearea = dynamic(import("components/Codearea"), { ssr: false })

export default class NewTemplate extends Component {
    static async getInitialProps(ctx) {
        if (redirectIfNotAuthenticated(ctx)) {
            return {}
        }
    }

    render() {
        const ancestors = [{ title: "Template list", href: "/templates" }]
        const initialValues = {
            title: "",
            isActive: false,
            content: ""
        }
        return (
            <Layout>
                <TemplateForm
                    type="create"
                    title="New Template"
                    values={initialValues}
                    ancestors={ancestors}
                />
            </Layout>
        )
    }
}
