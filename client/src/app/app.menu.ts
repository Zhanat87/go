export const APP_MENU = [
    {
        path: '',
        children: [
            {
                path: 'index',
                data: {
                    menu: {
                        title: 'Home',
                        icon: 'ion-android-home',
                        selected: false,
                        expanded: false,
                        order: 0
                    }
                }
            },
            {
                path: 'chat',
                data: {
                    menu: {
                        title: 'Chat',
                        icon: 'ion-ios-chatbubble-outline',
                        selected: false,
                        expanded: false,
                        order: 1
                    }
                }
            },
            {
                path: 'users',
                data: {
                    menu: {
                        title: 'Users',
                        icon: 'ion-stats-bars',
                        selected: false,
                        expanded: false,
                        order: 2
                    }
                }
            },
            {
                path: 'artists',
                data: {
                    menu: {
                        title: 'Artists',
                        icon: 'ion-ios-world',
                        selected: false,
                        expanded: false,
                        order: 3
                    }
                }
            },
            {
                path: 'albums',
                data: {
                    menu: {
                        title: 'Albums',
                        icon: 'ion-ios-pricetags-outline',
                        selected: false,
                        expanded: false,
                        order: 4
                    }
                }
            },
            {
                path: 'categories',
                data: {
                    menu: {
                        title: 'Categories',
                        icon: 'ion-grid',
                        selected: false,
                        expanded: false,
                        order: 5
                    }
                }
            },
            {
                path: 'partition-news',
                data: {
                    menu: {
                        title: 'Partitions news',
                        icon: 'ion-ios-bookmarks-outline',
                        selected: false,
                        expanded: false,
                        order: 6
                    }
                }
            },
            {
                path: 'shard-news',
                data: {
                    menu: {
                        title: 'Shards news',
                        icon: 'ion-ios-bookmarks',
                        selected: false,
                        expanded: false,
                        order: 7
                    }
                }
            },
            {
                path: 'replication-news',
                data: {
                    menu: {
                        title: 'Replication news',
                        icon: 'ion-clipboard',
                        selected: false,
                        expanded: false,
                        order: 8
                    }
                }
            },
        ]
    }
];
