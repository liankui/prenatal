/// the user package
package user {
    version: '0.1.0'
    authors: [{
        author: 'Eric'
        email:  'eric@1.q'
        organization: 'prenatal'
    }]

    dependencies: {
        'mojo.core': {repository: 'github.com/mojo-lang/core', version: '0.0.0-20211123010202-03f9f6e22fd2'}
        'mojo.geom': {repository: 'github.com/mojo-lang/geom', version: '0.0.0-20211121061729-3cc68d7658b2'}
    }
    repository: 'https://github.com/liankui/prenatal-server'
}
