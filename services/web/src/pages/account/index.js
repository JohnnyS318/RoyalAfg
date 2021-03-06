import React from "react";
import Layout from "../../components/layout";
import Front from "../../components/layout/front";
import ActionMenu from "../../components/actionMenu";
import ActionMenuLink from "../../components/actionMenu/link";
import { useRouter } from "next/router";
import { useSession } from "../../hooks/auth";

const Account = () => {
    const [session, loading] = useSession();
    const router = useRouter();

    if (loading) return null;

    if (!loading && !session) {
        router.push("/login");
        return <div>Access denied</div>;
    }

    return (
        <Layout>
            <div>
                <Front>{"Your Account " + session.user.name}</Front>
                <div className="px-10 pb-10 bg-gray-200">
                    <ActionMenu>
                        <ActionMenuLink href="/wallet">My Wallet</ActionMenuLink>
                    </ActionMenu>
                </div>
                {session && session.user ? session.user.name : <></>}
            </div>
        </Layout>
    );
};

export default Account;
