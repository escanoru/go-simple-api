pipeline {
    agent { label 'ansible' }
    options {
        buildDiscarder(logRotator(numToKeepStr: '7'))
    }
    environment {
        DOCKER_CREDENTIALS=credentials('docker-push')
        CONT_NAME = "${params.APP_NAME}"
    }
    parameters {
            string(name: 'APP_NAME', defaultValue: 'go-simple-api', description: 'Name of the app to push to Dockerhub')
            string(name: 'DOCKER_CREDENTIALS', defaultValue: 'docker-push', description: 'Name of the docker credentials on Jenkins')
    }
    stages {
        stage('Build') {
            steps {
                sh 'podman build -t ${DOCKER_CREDENTIALS_USR}/${CONT_NAME}:latest .'
            }
        }
        // stage('Login') {
        //     steps {
        //         sh 'echo ${env.DOCKER_CREDENTIALS_PSW} | docker login -u ${env.DOCKER_CREDENTIALS_USR} --password-stdin'
        //     }
        // }
        // stage('Push') {
        //     steps {
        //         sh 'docker push ${env.DOCKER_CREDENTIALS_USR}/${params.APP_NAME}:latest'
        //     }
        // }
    }
    // post {
    //     always {
    //         sh 'docker logout'
    //     }
    // }
}

