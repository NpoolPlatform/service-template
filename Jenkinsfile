pipeline {
  agent any
  environment {
    GOPROXY = 'https://goproxy.cn,direct'
  }
  tools {
    go 'go'
  }
  stages {
    stage('Clone') {
      steps {
        git(url: scm.userRemoteConfigs[0].url, branch: '$BRANCH_NAME', changelog: true, credentialsId: 'KK-github-key', poll: true)
      }
    }

    stage('Prepare') {
      steps {
        // Get linter and other build tools.
        sh 'go get -u golang.org/x/lint/golint'
        sh 'go get github.com/tebeka/go2xunit'
        sh 'go get github.com/t-yuki/gocover-cobertura'

        // Get dependencies
        sh 'go get golang.org/x/image/tiff/lzw'
        sh 'go get github.com/boombuler/barcode'
        sh 'make deps'
      }
    }

    stage('Linting') {
      when {
        expression { BUILD_TARGET == 'true' }
      }
      steps {
        sh 'make verify'
      }
    }

    stage('Compile') {
      when {
        expression { BUILD_TARGET == 'true' }
      }
      steps {
        sh 'make verify-build'
      }
    }

    stage('Unit Tests') {
      when {
        expression { BUILD_TARGET == 'true' }
      }
      steps {
        sh 'rm .apollo-base-config -rf'
        sh 'git clone https://github.com/NpoolPlatform/apollo-base-config.git .apollo-base-config'
        sh (returnStdout: false, script: '''
          devboxpod=`kubectl get pods -A | grep development-box | awk '{print $2}'`
          servicename="sample-service"

          PASSWORD=`kubectl get secret --namespace "kube-system" mysql-password-secret -o jsonpath="{.data.rootpassword}" | base64 --decode`
          kubectl -n kube-system exec mysql-0 -- mysql -h 127.0.0.1 -uroot -p$PASSWORD -P3306 -e "create database if not exists service_sample;"

          kubectl exec --namespace kube-system $devboxpod -- make -C /tmp/$servicename after-test || true
          kubectl exec --namespace kube-system $devboxpod -- rm -rf /tmp/$servicename || true
          kubectl cp ./ kube-system/$devboxpod:/tmp/$servicename

          username=`helm status rabbitmq --namespace kube-system | grep Username | awk -F ' : ' '{print $2}' | sed 's/"//g'`
          for vhost in `cat cmd/*/*.viper.yaml | grep hostname | awk '{print $2}' | sed 's/"//g' | sed 's/\\./-/g'`; do
            kubectl exec -it --namespace kube-system rabbitmq-0 -- rabbitmqctl add_vhost $vhost
            kubectl exec -it --namespace kube-system rabbitmq-0 -- rabbitmqctl set_permissions -p $vhost $username ".*" ".*" ".*"
            cd .apollo-base-config
            ./apollo-base-config.sh $APP_ID $TARGET_ENV $vhost
            ./apollo-item-config.sh $APP_ID $TARGET_ENV $vhost database_name service_sample
          done

          kubectl exec --namespace kube-system $devboxpod -- make -C /tmp/$servicename deps before-test test after-test
          kubectl exec --namespace kube-system $devboxpod -- rm -rf /tmp/$servicename
        '''.stripIndent())
      }
    }

    stage('Generate docker image') {
      when {
        expression { BUILD_TARGET == 'true' }
      }
      steps {
        sh 'make generate-docker-images'
      }
    }

    stage('Release docker image') {
      when {
        expression { RELEASE_TARGET == 'true' }
      }
      steps {
        sh 'make release-docker-images'
      }
    }

    stage('Deploy') {
      when {
        expression { DEPLOY_TARGET == 'true' }
      }
      steps {
        sh 'rm .apollo-base-config -rf'
        sh 'git clone https://github.com/NpoolPlatform/apollo-base-config.git .apollo-base-config'
        sh 'make deploy-to-k8s-cluster'
        sh (returnStdout: false, script: '''
          PASSWORD=`kubectl get secret --namespace "kube-system" mysql-password-secret -o jsonpath="{.data.rootpassword}" | base64 --decode`
          kubectl -n kube-system exec mysql-0 -- mysql -h 127.0.0.1 -uroot -p$PASSWORD -P3306 -e "create database if not exists service_sample;"
          username=`helm status rabbitmq --namespace kube-system | grep Username | awk -F ' : ' '{print $2}' | sed 's/"//g'`
          for vhost in `cat cmd/*/*.viper.yaml | grep hostname | awk '{print $2}' | sed 's/"//g' | sed 's/\\./-/g'`; do
            kubectl exec -it --namespace kube-system rabbitmq-0 -- rabbitmqctl add_vhost $vhost
            kubectl exec -it --namespace kube-system rabbitmq-0 -- rabbitmqctl set_permissions -p $vhost $username ".*" ".*" ".*"
            cd .apollo-base-config
            ./apollo-base-config.sh $APP_ID $TARGET_ENV $vhost
            ./apollo-item-config.sh $APP_ID $TARGET_ENV $vhost database_name service_sample
          done
        '''.stripIndent())
      }
    }

    stage('Post') {
      steps {
        // Assemble vet and lint info.
        // warnings parserConfigurations: [
        //   [pattern: 'govet.txt', parserName: 'Go Vet'],
        //   [pattern: 'golint.txt', parserName: 'Go Lint']
        // ]

        // sh 'go2xunit -fail -input gotest.txt -output gotest.xml'
        // junit "gotest.xml"
        sh 'echo Posting'
      }
    }
  }
  post('Report') {
    fixed {
      script {
        sh(script: 'bash $JENKINS_HOME/wechat-templates/send_wxmsg.sh fixed')
     }
      script {
        // env.ForEmailPlugin = env.WORKSPACE
        emailext attachmentsPattern: 'TestResults\\*.trx',
        body: '${FILE,path="$JENKINS_HOME/email-templates/success_email_tmp.html"}',
        mimeType: 'text/html',
        subject: currentBuild.currentResult + " : " + env.JOB_NAME,
        to: '$DEFAULT_RECIPIENTS'
      }
     }
    success {
      script {
        sh(script: 'bash $JENKINS_HOME/wechat-templates/send_wxmsg.sh successful')
     }
      script {
        // env.ForEmailPlugin = env.WORKSPACE
        emailext attachmentsPattern: 'TestResults\\*.trx',
        body: '${FILE,path="$JENKINS_HOME/email-templates/success_email_tmp.html"}',
        mimeType: 'text/html',
        subject: currentBuild.currentResult + " : " + env.JOB_NAME,
        to: '$DEFAULT_RECIPIENTS'
      }
     }
    failure {
      script {
        sh(script: 'bash $JENKINS_HOME/wechat-templates/send_wxmsg.sh failure')
     }
      script {
        // env.ForEmailPlugin = env.WORKSPACE
        emailext attachmentsPattern: 'TestResults\\*.trx',
        body: '${FILE,path="$JENKINS_HOME/email-templates/fail_email_tmp.html"}',
        mimeType: 'text/html',
        subject: currentBuild.currentResult + " : " + env.JOB_NAME,
        to: '$DEFAULT_RECIPIENTS'
      }
     }
    aborted {
      script {
        sh(script: 'bash $JENKINS_HOME/wechat-templates/send_wxmsg.sh aborted')
     }
      script {
        // env.ForEmailPlugin = env.WORKSPACE
        emailext attachmentsPattern: 'TestResults\\*.trx',
        body: '${FILE,path="$JENKINS_HOME/email-templates/fail_email_tmp.html"}',
        mimeType: 'text/html',
        subject: currentBuild.currentResult + " : " + env.JOB_NAME,
        to: '$DEFAULT_RECIPIENTS'
      }
     }
  }
}
