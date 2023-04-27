pipeline {
    agent {label 'worker-1'}
    stages{
        stage('Checkout') {
            steps {
                sshagent(credentials: ['WebApplication-Demo']) {
                    git url: 'https://github.com/EBambi/WebApplication-DEMO.git', branch: 'master', credentialsId: 'WebApplication-Demo'
                }
            }
        }
        stage('Build') {
            steps {
                sh 'go build main.go'
            }
        }
        stage('Deploy') {
            steps {
                sh 'JENKINS_NODE_COOKIE=dontKillMe nohup ./main &'
            }
        }
    }
}