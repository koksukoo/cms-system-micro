import Link from 'next/link'
import Layout from '../../components/BaseLayout'
import Widget from '../../components/Widget'
import Button from '../../components/ActionButton'

export default () => (
    <Layout>
        <Widget title="Template list">
            <Widget.Actions>
                <Link href="/templates/new">
                    <Button>+ New Template</Button>
                </Link>
            </Widget.Actions>
            <Widget.Body>
            </Widget.Body>
        </Widget>
    </Layout>
)