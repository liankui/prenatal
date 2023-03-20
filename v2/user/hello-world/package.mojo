/// the hello_world package
package hello_world {
    version: 'v0.0.1'
    authors: [{
        author: ''
        email:  ''
        organization: 'mojo-lang.org'
    }]
    
    dependencies: {
        'mojo.core': {repository: 'github.com/mojo-lang/core', version: '0.0.0-20211123010202-03f9f6e22fd2'}
        'mojo.geom': {repository: 'github.com/mojo-lang/geom', version: '0.0.0-20211121061729-3cc68d7658b2'}
    }
    repository: 'git.company.com/examples/hello-world'
}
