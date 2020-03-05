//This can be used if the node does not have the given tool
// node {
//     // Install the desired Go version
//     def root = tool name: 'Go 1.4', type: 'go'

//     // Export environment variables pointing to the directory where Go was installed
//     withEnv(["GOROOT=${root}", "PATH+GO=${root}/bin"]) {
//         sh 'go version'
//     }
// }
pipeline {
	agent any
    tools {
        go 'Go'
    }
	stages {
		stage('Checkout') {
            steps{
				sh 'mkdir github.com'
				sh 'mkdir github.com/NehemiahG7'
				sh 'cd github.com/NehemiahG7'
			    checkout scm
            }
		}
		stage('Build') {
			steps {
				sh 'go build *.go'
			}
		}
	}
	post {
		always {
			sh 'cd ../..'
			deleteDir
		}
	}
}
