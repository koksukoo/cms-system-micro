import { rem } from 'utils/style';

const ProfileDropdown = () => (
    <nav>
        <button>
            <span>Hello, Mikko</span>
            <img className="pic" src="/static/img/profile-placeholder.jpg" />
        </button>

        <style jsx>{`
            button {
                background-color: transparent;
                border: 0;
                display: flex;
                height: 50px;
                font-size: ${rem(14)};
                align-items: center;
            }
            .pic {
                border-radius: 50%;
                margin-left: 20px;
                height: 40px;
                width: 40px;
            }
        `}</style>
    </nav>
)

export default ProfileDropdown;