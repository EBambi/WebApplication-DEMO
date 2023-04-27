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
        stage('Connect to EC2') {
            steps {
                sshagent(['/home/jenkins2/id_rsa']) {
                    sshCommand remote: ubuntu@ec2-13-58-91-243.us-east-2.compute.amazonaws.com, command: 'echo "Connection Successfully"'
                }
            }
        }
        /*stage('Deploy') {
            steps {
                sh ''
                sh 'JENKINS_NODE_COOKIE=dontKillMe nohup ./main &'
            }
        }*/
    }
}