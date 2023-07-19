


## 注意事项 
1.由于 kettle 对应的jar包 在中央仓库有但是无法下载下来完整jar 因此 直接 官网下载对应的PDI（Pentaho Data Integration (Base Install)）

官网：https://www.hitachivantara.com/en-us/products/pentaho-platform/data-integration-analytics/pentaho-community-edition.html

2.下载完成后 到pdi-ce-9.4.0.0-343\data-integration\lib文件夹找到 kettle-core和kettle-engine jar文件然后maven本地引入即可

3.启动初始化过程中还会遇到错误 看下面maven依赖

4. resource 下面增加 xml文件<kettle-password-encoder-plugins.xml>

```
<!--解决报错：org.pentaho.di.core.exception.KettleException:
Unable to find plugin with ID 'Kettle'.  If this is a test, make sure kettle-core tests jar is a dependency.  If this is live make sure a kettle-password-encoder-plugins.xml exits in the classpath
-->

<password-encoder-plugins>
    <password-encoder-plugin id="Kettle">
        <description>Kettle Password Encoder</description>
        <classname>org.pentaho.di.core.encryption.KettleTwoWayPasswordEncoder</classname>
    </password-encoder-plugin>
</password-encoder-plugins>
```

## 本地引入方式

```
<!--   kettle -core 核心     -->
        <dependency>
            <groupId>org.pentaho</groupId>
            <artifactId>kettle-engine</artifactId>
            <scope>system</scope>
            <version>9.4.0.0-343</version>
            <systemPath>${basedir}/libs/kettle-engine-9.4.0.0-343.jar </systemPath>
        </dependency>
        <dependency>
            <groupId>org.pentaho</groupId>
            <artifactId>kettle-core</artifactId>
            <scope>system</scope>
            <version>9.4.0.0-343</version>
            <systemPath>${basedir}/libs/kettle-core-9.4.0.0-343.jar </systemPath>
        </dependency>

        <!-- java.lang.NoClassDefFoundError: org/apache/commons/vfs2/FileSelector-->
        <dependency>
            <groupId>org.apache.commons</groupId>
            <artifactId>commons-vfs2</artifactId>
            <version>2.8.0</version>
        </dependency>

<!--   java.lang.NoClassDefFoundError: com/google/common/util/concurrent/SettableFuture     -->
        <dependency>
            <groupId>kettle.guava</groupId>
            <artifactId>guava</artifactId>
            <version>17.0</version>
            <scope>system</scope>
            <systemPath>${basedir}/libs/guava-17.0.jar </systemPath>
        </dependency>

        <!--java.lang.NoClassDefFoundError: org/apache/commons/io/IOUtils-->
        <dependency>
            <groupId>commons-io</groupId>
            <artifactId>commons-io</artifactId>
            <version>2.4</version>
        </dependency>

        <!--NoClassDefFoundError: org/apache/commons/lang/builder/HashCodeBuilder-->
        <!-- https://mvnrepository.com/artifact/commons-lang/commons-lang -->
        <dependency>
            <groupId>commons-lang</groupId>
            <artifactId>commons-lang</artifactId>
            <version>2.6</version>
        </dependency>

        <!--   resource 下增加 kettle-password 报错 Type org.pentaho.di.core.encryption.TwoWayPasswordEncoderInterface not present-->
        <dependency>
            <groupId>org.pentaho</groupId>
            <artifactId>pentaho-encryption-support</artifactId>
            <version>9.4.0.0-343</version>
            <scope>system</scope>
            <systemPath>${basedir}/libs/pentaho-encryption-support-9.4.0.0-343.jar </systemPath>
        </dependency>

         <!--     java:   org.pentaho.metastore.api.IMetaStore   -->
        <dependency>
            <groupId>org.pentaho</groupId>
            <artifactId>metastore</artifactId>
            <version>9.4.0.0-343</version>
            <scope>system</scope>
            <systemPath>${basedir}/libs/metastore-9.4.0.0-343.jar </systemPath>
        </dependency>


         <!--ClassNotFoundException: org.apache.http.client.methods.HttpPost-->
        <dependency>
            <groupId>org.apache.httpcomponents</groupId>
            <artifactId>httpclient</artifactId>
            <version>4.3.2</version>
        </dependency>

         <!--java.lang.NoClassDefFoundError: org/owasp/encoder/Encode-->
        <dependency>
            <groupId>org.owasp.encoder</groupId>
            <artifactId>encoder-esapi</artifactId>
            <version>1.2.1</version>
        </dependency>

```