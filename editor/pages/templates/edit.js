import { Component } from 'react'
import SaveIcon from 'react-icons/lib/fa/floppy-o'
import dynamic from 'next/dynamic'
import Router, { withRouter } from 'next/router'
import Layout from 'components/BaseLayout'
import Widget from 'components/Widget'
import Button from 'components/ActionButton'
import Input from 'components/Input'
import 'codemirror/lib/codemirror.css'
import { updateTemplate, fetchTemplate } from 'io'

const Codearea = dynamic(import('components/Codearea'), { ssr: false })

class EditTemplate extends Component {
    constructor(props) {
        super(props)
        this.handleForm = this.handleForm.bind(this)
    }

    static async getInitialProps(ctx) {
        const { id } = ctx.query
        const res = await fetchTemplate(id)
        const data = await res.json()

        return await {
            template: data
        }
    }

    async handleForm() {
        const { title, isActive, content } = this.form
        const { id } = this.props.router.query
        const data = {
            title: title.value,
            isActive: isActive.checked,
            content: content.value,
        }

        const res = updateTemplate(id, data)
        res.then(() => {
            Router.push('/templates')
        })
    }

    render() {
        const { template } = this.props
        const ancestors = [
            { title: 'Template list', href: '/templates' },
        ];

        return (
            <Layout>
                <Widget title="Edit Template" ancestors={ancestors}>
                    <Widget.Actions>
                        <Button onClick={this.handleForm}><SaveIcon /> Save template</Button>
                    </Widget.Actions>
                    <Widget.Body>
                        <form onSubmit={updateTemplate} ref={node => (this.form = node)}>
                            <Input name="title" label="Title" initialValue={template.title} required />
                            <Input name="isActive" label="Is active" initialChecked={template.isActive} type="checkbox" />
                            <Codearea name="content" label="Content" initialValue={template.content} rows="20" />
                        </form>
                    </Widget.Body>
                </Widget>
                <style jsx>{`
                    form {
                        margin: 30px 0;
                    }
                `}</style>
            </Layout>
        )
    }
}

export default withRouter(EditTemplate)