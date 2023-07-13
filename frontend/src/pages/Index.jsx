import Sidebar from "../components/Sidebar";
import { Outlet } from "react-router-dom";

const Index = () => {
    return (<>
        <div className="home flex gap-8">
            <div className="sidebar">
                <Sidebar />
            </div>

            <div className="content">
                <main>
                    <Outlet />
                </main>
            </div>
        </div>

    </>);
}

export default Index;