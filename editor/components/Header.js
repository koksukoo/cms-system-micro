import ProfileDropdown from './ProfileDropdown'

const Header = () => (
    <header>
        <h1>
            <strong>CMS</strong> Editor
        </h1>

        <ProfileDropdown />
        
        <style jsx>{`
            background-color: #fff;
            display: flex;
            justify-content: space-between;
            align-items: center;
            padding: 0 15px;

            strong {
                padding: 0;
            }

            h1 {
                font-weight: 300;
                padding: 0; 
            }
        `}
        </style>
    </header>
)

export default Header