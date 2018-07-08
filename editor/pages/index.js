import Link from 'next/link'
import PlusIcon from 'react-icons/lib/md/add'
import Layout from '../components/BaseLayout'
import Widget from '../components/Widget'
import HierarchyBody from '../components/widgets/Hierarchy'
import { fetchPageHierarchy } from '../io';
import Button from '../components/ActionButton'

const Index = (props) => (
    <Layout>
        <Widget title="Site map">
            <Widget.Actions>
                <Link href="/pages/new">
                    <Button><PlusIcon /> New Page</Button>
                </Link>
            </Widget.Actions>

            <Widget.Body>
                <HierarchyBody children={props.hierarchy} />
            </Widget.Body>
        </Widget>
    </Layout>
)

Index.getInitialProps = async () => {
    const res = await fetchPageHierarchy()
    const data = await res.json()
    
    return await {
        hierarchy: data,
    }
}

export default Index