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
        /*stage('Connect to EC2') {
            steps {
                sshagent(['AWS-Credential']) {
                    sshCommand remote: ec2-ubuntu@ec2-12-58-91-243.compute-1.amazonaws.com, command: 'echo "Connection Successfully"'
                }
            }
        }*/
        stage('Deploy') {
            steps {
                sh 'scp -i "/home/jenkins2/id_rsa" main ubuntu@ec2-13-58-91-243.us-east-2.compute.amazonaws.com:/home/ubuntu/'
                sh 'ssh -i "/home/jenkins2/id_rsa" ubuntu@ec2-13-58-91-243.us-east-2.compute.amazonaws.com'
                sh 'JENKINS_NODE_COOKIE=dontKillMe nohup ./main &'
            }
        }
    }
}