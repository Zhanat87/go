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
                path: 'users',
                data: {
                    menu: {
                        title: 'Users',
                        icon: 'ion-stats-bars',
                        selected: false,
                        expanded: false,
                        order: 1
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
                        order: 2
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
                        order: 3
                    }
                }
            },
        ]
    }
];
