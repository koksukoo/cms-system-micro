import { Component } from 'react'
import SaveIcon from 'react-icons/lib/fa/floppy-o'
import dynamic from 'next/dynamic'
import Router from 'next/router'
import Layout from '~/components/BaseLayout'
import Widget from '~/components/Widget'
import Button from '~/components/ActionButton'
import Input from '~/components/Input'
import { createTemplate } from '~/io'

const Codearea = dynamic(import('~/components/Codearea'), { ssr: false })

export default class NewTemplate extends Component {
    constructor(props) {
        super(props)
        this.handleForm = this.handleForm.bind(this)
    }

    async handleForm() {
        const {title, isActive, content} = this.form
        const data = {
            title: title.value,
            isActive: isActive.checked,
            content: content.value,
        }

        const res = createTemplate(data)
        res.then(() => {
            Router.push('/templates')
        })
    }

    render() {
        const ancestors = [
            { title: 'Template list', href: '/templates' },
        ];
        return (
            <Layout>
                <Widget title="New Template" ancestors={ancestors}>
                    <Widget.Actions>
                        <Button onClick={this.handleForm}><SaveIcon /> Save template</Button>
                    </Widget.Actions>
                    <Widget.Body>
                        <form onSubmit={createTemplate} ref={node => (this.form = node)}>
                            <Input name="title" label="Title" required />
                            <Input name="isActive" label="Is active" type="checkbox" />
                            <Codearea name="content" label="Content" rows="20" />
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