const Sidebar = () => {
    return (<>
        <div className="sidebar">
            <div className="content flex flex-col gap-4 bg-white shadow-lg rounded-md px-8 py-3">
                <div className="all-post">
                    <a href="/posts">All Post</a>
                </div>
                <hr />
                <div className="all-post">
                    <a href="/new">Add New</a>
                </div>
                <hr />
                <div className="all-post">
                    <a href="/preview">Preview</a>
                </div>
            </div>
        </div>



    </>);
}

export default Sidebar;