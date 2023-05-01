pipeline {
    agent { label 'ansible' }
    options {
        buildDiscarder(logRotator(numToKeepStr: '7'))
    }
    environment {
        DOCKER_CREDS = credentials("${params.DOCKER_CREDENTIALS}")
    }
    parameters {
            string(name: 'APP_NAME', defaultValue: 'empty', description: 'Name of the app to push to Dockerhub')
            string(name: 'DOCKER_CREDENTIALS', defaultValue: 'empty', description: 'Name of the docker credentials on Jenkins')
    }
    stages {
        stage('Build') {
            steps {
                sh 'docker build -t ${DOCKER_CREDS_USR}/${params.APP_NAME}:latest .'
            }
        }
        stage('Login') {
            steps {
                sh 'echo ${DOCKER_CREDS_PSW} | docker login -u ${DOCKER_CREDS_USR} --password-stdin'
            }
        }
        stage('Push') {
            steps {
                sh 'docker push ${DOCKER_CREDS_USR}/${APP_NAME}:latest'
            }
        }
    }
    // post {
    //     always {
    //         sh 'docker logout'
    //     }
    // }
}

