import { Component } from "react"
import dynamic from "next/dynamic"
import Router, { withRouter } from "next/router"
import Layout from "components/BaseLayout"
import TemplateForm from "components/TemplateForm"
import "codemirror/lib/codemirror.css"
import { fetchTemplate } from "io"

const Codearea = dynamic(import("components/Codearea"), { ssr: false })

class EditTemplate extends Component {
    static async getInitialProps(ctx) {
        const { id } = ctx.query
        const res = await fetchTemplate(id)
        const data = await res.json()

        return await {
            initialValues: data
        }
    }

    render() {
        const { initialValues } = this.props
        const ancestors = [{ title: "Template list", href: "/templates" }]
        return (
            <Layout>
                <TemplateForm
                    type="edit"
                    title="Edit Template"
                    values={initialValues}
                    ancestors={ancestors}
                />
            </Layout>
        )
    }
}

export default withRouter(EditTemplate)
