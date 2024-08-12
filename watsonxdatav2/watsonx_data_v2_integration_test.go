//go:build integration

/**
 * (C) Copyright IBM Corp. 2024.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *      http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package watsonxdatav2_test

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/IBM/go-sdk-core/v5/core"
	"github.com/IBM/watsonxdata-go-sdk/watsonxdatav2"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

/**
 * This file contains an integration test for the watsonxdatav2 package.
 *
 * Notes:
 *
 * The integration test will automatically skip tests if the required config file is not available.
 */

var _ = Describe(`WatsonxDataV2 Integration Tests`, func() {
	const externalConfigFile = "../watsonx_data_v2.env"

	var (
		err          error
		watsonxDataService *watsonxdatav2.WatsonxDataV2
		serviceURL   string
		config       map[string]string
	)

	var shouldSkipTest = func() {
		Skip("External configuration is not available, skipping tests...")
	}

	Describe(`External configuration`, func() {
		It("Successfully load the configuration", func() {
			_, err = os.Stat(externalConfigFile)
			if err != nil {
				Skip("External configuration file not found, skipping tests: " + err.Error())
			}

			os.Setenv("IBM_CREDENTIALS_FILE", externalConfigFile)
			config, err = core.GetServiceProperties(watsonxdatav2.DefaultServiceName)
			if err != nil {
				Skip("Error loading service properties, skipping tests: " + err.Error())
			}
			serviceURL = config["URL"]
			if serviceURL == "" {
				Skip("Unable to load service URL configuration property, skipping tests")
			}

			fmt.Fprintf(GinkgoWriter, "Service URL: %v\n", serviceURL)
			shouldSkipTest = func() {}
		})
	})

	Describe(`Client initialization`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It("Successfully construct the service client instance", func() {
			watsonxDataServiceOptions := &watsonxdatav2.WatsonxDataV2Options{}

			watsonxDataService, err = watsonxdatav2.NewWatsonxDataV2UsingExternalConfig(watsonxDataServiceOptions)
			Expect(err).To(BeNil())
			Expect(watsonxDataService).ToNot(BeNil())
			Expect(watsonxDataService.Service.Options.URL).To(Equal(serviceURL))

			core.SetLogger(core.NewLogger(core.LevelDebug, log.New(GinkgoWriter, "", log.LstdFlags), log.New(GinkgoWriter, "", log.LstdFlags)))
			watsonxDataService.EnableRetries(4, 30*time.Second)
		})
	})

	Describe(`ListBucketRegistrations - Get bucket registrations`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`ListBucketRegistrations(listBucketRegistrationsOptions *ListBucketRegistrationsOptions)`, func() {
			listBucketRegistrationsOptions := &watsonxdatav2.ListBucketRegistrationsOptions{
				AuthInstanceID: core.StringPtr("testString"),
			}

			bucketRegistrationCollection, response, err := watsonxDataService.ListBucketRegistrations(listBucketRegistrationsOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(bucketRegistrationCollection).ToNot(BeNil())
		})
	})

	Describe(`CreateBucketRegistration - Register bucket`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`CreateBucketRegistration(createBucketRegistrationOptions *CreateBucketRegistrationOptions)`, func() {
			bucketDetailsModel := &watsonxdatav2.BucketDetails{
				AccessKey: core.StringPtr("b9cbf248ea5c4c96947e64407108559j"),
				BucketName: core.StringPtr("sample-bucket"),
				Endpoint: core.StringPtr("https://s3.<region>.cloud-object-storage.appdomain.cloud/"),
				SecretKey: core.StringPtr("13b4045cac1a0be54c9fjbe53cb22df5fn397cd2c45b66c87"),
			}

			bucketCatalogModel := &watsonxdatav2.BucketCatalog{
				CatalogName: core.StringPtr("sampleCatalog"),
				CatalogTags: []string{"catalog_tag_1", "catalog_tag_2"},
				CatalogType: core.StringPtr("iceberg"),
			}

			createBucketRegistrationOptions := &watsonxdatav2.CreateBucketRegistrationOptions{
				BucketDetails: bucketDetailsModel,
				BucketType: core.StringPtr("ibm_cos"),
				Description: core.StringPtr("COS bucket for customer data"),
				ManagedBy: core.StringPtr("ibm"),
				AssociatedCatalog: bucketCatalogModel,
				BucketDisplayName: core.StringPtr("sample-bucket-displayname"),
				Region: core.StringPtr("us-south"),
				Tags: []string{"bucket-tag1", "bucket-tag2"},
				AuthInstanceID: core.StringPtr("testString"),
			}

			bucketRegistration, response, err := watsonxDataService.CreateBucketRegistration(createBucketRegistrationOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(201))
			Expect(bucketRegistration).ToNot(BeNil())
		})
	})

	Describe(`GetBucketRegistration - Get bucket`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`GetBucketRegistration(getBucketRegistrationOptions *GetBucketRegistrationOptions)`, func() {
			getBucketRegistrationOptions := &watsonxdatav2.GetBucketRegistrationOptions{
				BucketID: core.StringPtr("testString"),
				AuthInstanceID: core.StringPtr("testString"),
			}

			bucketRegistration, response, err := watsonxDataService.GetBucketRegistration(getBucketRegistrationOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(bucketRegistration).ToNot(BeNil())
		})
	})

	Describe(`UpdateBucketRegistration - Update bucket`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`UpdateBucketRegistration(updateBucketRegistrationOptions *UpdateBucketRegistrationOptions)`, func() {
			bucketDetailsModel := &watsonxdatav2.BucketDetails{
				AccessKey: core.StringPtr("b9cbf248ea5c4c96947e64407108559j"),
				BucketName: core.StringPtr("sample-bucket"),
				Endpoint: core.StringPtr("https://s3.<region>.cloud-object-storage.appdomain.cloud/"),
				SecretKey: core.StringPtr("13b4045cac1a0be54c9fjbe53cb22df5fn397cd2c45b66c87"),
			}

			bucketRegistrationPatchModel := &watsonxdatav2.BucketRegistrationPatch{
				BucketDetails: bucketDetailsModel,
				BucketDisplayName: core.StringPtr("sample-bucket-displayname"),
				Description: core.StringPtr("COS bucket for customer data"),
				Tags: []string{"testbucket", "userbucket"},
			}
			bucketRegistrationPatchModelAsPatch, asPatchErr := bucketRegistrationPatchModel.AsPatch()
			Expect(asPatchErr).To(BeNil())

			updateBucketRegistrationOptions := &watsonxdatav2.UpdateBucketRegistrationOptions{
				BucketID: core.StringPtr("testString"),
				Body: bucketRegistrationPatchModelAsPatch,
				AuthInstanceID: core.StringPtr("testString"),
			}

			bucketRegistration, response, err := watsonxDataService.UpdateBucketRegistration(updateBucketRegistrationOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(bucketRegistration).ToNot(BeNil())
		})
	})

	Describe(`CreateActivateBucket - Activate Bucket`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`CreateActivateBucket(createActivateBucketOptions *CreateActivateBucketOptions)`, func() {
			createActivateBucketOptions := &watsonxdatav2.CreateActivateBucketOptions{
				BucketID: core.StringPtr("testString"),
				AuthInstanceID: core.StringPtr("testString"),
			}

			createActivateBucketCreatedBody, response, err := watsonxDataService.CreateActivateBucket(createActivateBucketOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(201))
			Expect(createActivateBucketCreatedBody).ToNot(BeNil())
		})
	})

	Describe(`ListBucketObjects - List bucket objects`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`ListBucketObjects(listBucketObjectsOptions *ListBucketObjectsOptions)`, func() {
			listBucketObjectsOptions := &watsonxdatav2.ListBucketObjectsOptions{
				BucketID: core.StringPtr("testString"),
				AuthInstanceID: core.StringPtr("testString"),
			}

			bucketRegistrationObjectCollection, response, err := watsonxDataService.ListBucketObjects(listBucketObjectsOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(bucketRegistrationObjectCollection).ToNot(BeNil())
		})
	})

	Describe(`ListDatabaseRegistrations - Get databases`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`ListDatabaseRegistrations(listDatabaseRegistrationsOptions *ListDatabaseRegistrationsOptions)`, func() {
			listDatabaseRegistrationsOptions := &watsonxdatav2.ListDatabaseRegistrationsOptions{
				AuthInstanceID: core.StringPtr("testString"),
			}

			databaseRegistrationCollection, response, err := watsonxDataService.ListDatabaseRegistrations(listDatabaseRegistrationsOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(databaseRegistrationCollection).ToNot(BeNil())
		})
	})

	Describe(`CreateDatabaseRegistration - Add/Create database`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`CreateDatabaseRegistration(createDatabaseRegistrationOptions *CreateDatabaseRegistrationOptions)`, func() {
			databaseCatalogModel := &watsonxdatav2.DatabaseCatalog{
				CatalogName: core.StringPtr("sampleCatalog"),
				CatalogTags: []string{"catalog_tag_1", "catalog_tag_2"},
				CatalogType: core.StringPtr("iceberg"),
			}

			databaseDetailsModel := &watsonxdatav2.DatabaseDetails{
				Certificate: core.StringPtr("contents of a pem/crt file"),
				CertificateExtension: core.StringPtr("pem/crt"),
				DatabaseName: core.StringPtr("new_database"),
				Hostname: core.StringPtr("db2@<hostname>.com"),
				HostnameInCertificate: core.StringPtr("samplehostname"),
				Hosts: core.StringPtr("abc.com:1234,xyz.com:4321"),
				Password: core.StringPtr("samplepassword"),
				Port: core.Int64Ptr(int64(4553)),
				Sasl: core.BoolPtr(true),
				Ssl: core.BoolPtr(true),
				Tables: core.StringPtr("kafka_table_name"),
				Username: core.StringPtr("sampleuser"),
				ValidateServerCertificate: core.BoolPtr(true),
			}

			databaseRegistrationPrototypeDatabasePropertiesItemsModel := &watsonxdatav2.DatabaseRegistrationPrototypeDatabasePropertiesItems{
				Encrypt: core.BoolPtr(true),
				Key: core.StringPtr("abc"),
				Value: core.StringPtr("xyz"),
			}

			createDatabaseRegistrationOptions := &watsonxdatav2.CreateDatabaseRegistrationOptions{
				DatabaseDisplayName: core.StringPtr("new_database"),
				DatabaseType: core.StringPtr("db2"),
				AssociatedCatalog: databaseCatalogModel,
				CreatedOn: core.StringPtr("1686792721"),
				DatabaseDetails: databaseDetailsModel,
				DatabaseProperties: []watsonxdatav2.DatabaseRegistrationPrototypeDatabasePropertiesItems{*databaseRegistrationPrototypeDatabasePropertiesItemsModel},
				Description: core.StringPtr("db2 extenal database description"),
				Tags: []string{"testdatabase", "userdatabase"},
				AuthInstanceID: core.StringPtr("testString"),
			}

			databaseRegistration, response, err := watsonxDataService.CreateDatabaseRegistration(createDatabaseRegistrationOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(201))
			Expect(databaseRegistration).ToNot(BeNil())
		})
	})

	Describe(`GetDatabase - Get database`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`GetDatabase(getDatabaseOptions *GetDatabaseOptions)`, func() {
			getDatabaseOptions := &watsonxdatav2.GetDatabaseOptions{
				DatabaseID: core.StringPtr("testString"),
				AuthInstanceID: core.StringPtr("testString"),
			}

			databaseRegistration, response, err := watsonxDataService.GetDatabase(getDatabaseOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(databaseRegistration).ToNot(BeNil())
		})
	})

	Describe(`UpdateDatabase - Update database`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`UpdateDatabase(updateDatabaseOptions *UpdateDatabaseOptions)`, func() {
			databaseRegistrationPatchDatabaseDetailsModel := &watsonxdatav2.DatabaseRegistrationPatchDatabaseDetails{
				Password: core.StringPtr("samplepassword"),
				Username: core.StringPtr("sampleuser"),
			}

			databaseRegistrationPatchModel := &watsonxdatav2.DatabaseRegistrationPatch{
				DatabaseDetails: databaseRegistrationPatchDatabaseDetailsModel,
				DatabaseDisplayName: core.StringPtr("new_database"),
				Description: core.StringPtr("External database description"),
				Tags: []string{"testdatabase", "userdatabase"},
			}
			databaseRegistrationPatchModelAsPatch, asPatchErr := databaseRegistrationPatchModel.AsPatch()
			Expect(asPatchErr).To(BeNil())

			updateDatabaseOptions := &watsonxdatav2.UpdateDatabaseOptions{
				DatabaseID: core.StringPtr("testString"),
				Body: databaseRegistrationPatchModelAsPatch,
				AuthInstanceID: core.StringPtr("testString"),
			}

			databaseRegistration, response, err := watsonxDataService.UpdateDatabase(updateDatabaseOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(databaseRegistration).ToNot(BeNil())
		})
	})

	Describe(`ListOtherEngines - List other engines`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`ListOtherEngines(listOtherEnginesOptions *ListOtherEnginesOptions)`, func() {
			listOtherEnginesOptions := &watsonxdatav2.ListOtherEnginesOptions{
				AuthInstanceID: core.StringPtr("testString"),
			}

			otherEngineCollection, response, err := watsonxDataService.ListOtherEngines(listOtherEnginesOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(otherEngineCollection).ToNot(BeNil())
		})
	})

	Describe(`CreateOtherEngine - Create other engine`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`CreateOtherEngine(createOtherEngineOptions *CreateOtherEngineOptions)`, func() {
			otherEngineDetailsBodyModel := &watsonxdatav2.OtherEngineDetailsBody{
				ConnectionString: core.StringPtr("1.2.3.4"),
				EngineType: core.StringPtr("netezza"),
			}

			createOtherEngineOptions := &watsonxdatav2.CreateOtherEngineOptions{
				EngineDetails: otherEngineDetailsBodyModel,
				EngineDisplayName: core.StringPtr("sampleEngine01"),
				Description: core.StringPtr("external engine description"),
				Origin: core.StringPtr("external"),
				Tags: []string{"tag1", "tag2"},
				Type: core.StringPtr("netezza"),
				AuthInstanceID: core.StringPtr("testString"),
			}

			otherEngine, response, err := watsonxDataService.CreateOtherEngine(createOtherEngineOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(201))
			Expect(otherEngine).ToNot(BeNil())
		})
	})

	Describe(`ListDb2Engines - Get list of db2 engines`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`ListDb2Engines(listDb2EnginesOptions *ListDb2EnginesOptions)`, func() {
			listDb2EnginesOptions := &watsonxdatav2.ListDb2EnginesOptions{
				AuthInstanceID: core.StringPtr("testString"),
			}

			db2EngineCollection, response, err := watsonxDataService.ListDb2Engines(listDb2EnginesOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(db2EngineCollection).ToNot(BeNil())
		})
	})

	Describe(`CreateDb2Engine - Create db2 engine`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`CreateDb2Engine(createDb2EngineOptions *CreateDb2EngineOptions)`, func() {
			db2EngineDetailsBodyModel := &watsonxdatav2.Db2EngineDetailsBody{
				ConnectionString: core.StringPtr("1.2.3.4"),
			}

			createDb2EngineOptions := &watsonxdatav2.CreateDb2EngineOptions{
				Origin: core.StringPtr("external"),
				Description: core.StringPtr("db2 engine description"),
				EngineDetails: db2EngineDetailsBodyModel,
				EngineDisplayName: core.StringPtr("sampleEngine"),
				Tags: []string{"tag1", "tag2"},
				AuthInstanceID: core.StringPtr("testString"),
			}

			db2Engine, response, err := watsonxDataService.CreateDb2Engine(createDb2EngineOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(201))
			Expect(db2Engine).ToNot(BeNil())
		})
	})

	Describe(`UpdateDb2Engine - Update db2 engine`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`UpdateDb2Engine(updateDb2EngineOptions *UpdateDb2EngineOptions)`, func() {
			db2EnginePatchModel := &watsonxdatav2.Db2EnginePatch{
				Description: core.StringPtr("db2 engine updated description"),
				EngineDisplayName: core.StringPtr("sampleEngine"),
				Tags: []string{"tag1", "tag2"},
			}
			db2EnginePatchModelAsPatch, asPatchErr := db2EnginePatchModel.AsPatch()
			Expect(asPatchErr).To(BeNil())

			updateDb2EngineOptions := &watsonxdatav2.UpdateDb2EngineOptions{
				EngineID: core.StringPtr("testString"),
				Body: db2EnginePatchModelAsPatch,
				AuthInstanceID: core.StringPtr("testString"),
			}

			db2Engine, response, err := watsonxDataService.UpdateDb2Engine(updateDb2EngineOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(db2Engine).ToNot(BeNil())
		})
	})

	Describe(`ListNetezzaEngines - Get list of netezza engines`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`ListNetezzaEngines(listNetezzaEnginesOptions *ListNetezzaEnginesOptions)`, func() {
			listNetezzaEnginesOptions := &watsonxdatav2.ListNetezzaEnginesOptions{
				AuthInstanceID: core.StringPtr("testString"),
			}

			netezzaEngineCollection, response, err := watsonxDataService.ListNetezzaEngines(listNetezzaEnginesOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(netezzaEngineCollection).ToNot(BeNil())
		})
	})

	Describe(`CreateNetezzaEngine - Create netezza engine`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`CreateNetezzaEngine(createNetezzaEngineOptions *CreateNetezzaEngineOptions)`, func() {
			netezzaEngineDetailsBodyModel := &watsonxdatav2.NetezzaEngineDetailsBody{
				ConnectionString: core.StringPtr("1.2.3.4"),
			}

			createNetezzaEngineOptions := &watsonxdatav2.CreateNetezzaEngineOptions{
				Origin: core.StringPtr("external"),
				Description: core.StringPtr("netezza engine description"),
				EngineDetails: netezzaEngineDetailsBodyModel,
				EngineDisplayName: core.StringPtr("sampleEngine"),
				Tags: []string{"tag1", "tag2"},
				AuthInstanceID: core.StringPtr("testString"),
			}

			netezzaEngine, response, err := watsonxDataService.CreateNetezzaEngine(createNetezzaEngineOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(201))
			Expect(netezzaEngine).ToNot(BeNil())
		})
	})

	Describe(`UpdateNetezzaEngine - Update netezza engine`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`UpdateNetezzaEngine(updateNetezzaEngineOptions *UpdateNetezzaEngineOptions)`, func() {
			netezzaEnginePatchModel := &watsonxdatav2.NetezzaEnginePatch{
				Description: core.StringPtr("netezza engine updated description"),
				EngineDisplayName: core.StringPtr("sampleEngine"),
				Tags: []string{"tag1", "tag2"},
			}
			netezzaEnginePatchModelAsPatch, asPatchErr := netezzaEnginePatchModel.AsPatch()
			Expect(asPatchErr).To(BeNil())

			updateNetezzaEngineOptions := &watsonxdatav2.UpdateNetezzaEngineOptions{
				EngineID: core.StringPtr("testString"),
				Body: netezzaEnginePatchModelAsPatch,
				AuthInstanceID: core.StringPtr("testString"),
			}

			netezzaEngine, response, err := watsonxDataService.UpdateNetezzaEngine(updateNetezzaEngineOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(netezzaEngine).ToNot(BeNil())
		})
	})

	Describe(`ListPrestissimoEngines - Get list of prestissimo engines`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`ListPrestissimoEngines(listPrestissimoEnginesOptions *ListPrestissimoEnginesOptions)`, func() {
			listPrestissimoEnginesOptions := &watsonxdatav2.ListPrestissimoEnginesOptions{
				AuthInstanceID: core.StringPtr("testString"),
			}

			prestissimoEngineCollection, response, err := watsonxDataService.ListPrestissimoEngines(listPrestissimoEnginesOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(prestissimoEngineCollection).ToNot(BeNil())
		})
	})

	Describe(`CreatePrestissimoEngine - Create prestissimo engine`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`CreatePrestissimoEngine(createPrestissimoEngineOptions *CreatePrestissimoEngineOptions)`, func() {
			prestissimoNodeDescriptionBodyModel := &watsonxdatav2.PrestissimoNodeDescriptionBody{
				NodeType: core.StringPtr("worker"),
				Quantity: core.Int64Ptr(int64(38)),
			}

			prestissimoEndpointsModel := &watsonxdatav2.PrestissimoEndpoints{
				ApplicationsApi: core.StringPtr("$HOST/v4/analytics_engines/c7b3fccf-badb-46b0-b1ef-9b3154424021/spark_applications/<application_id>"),
				HistoryServerEndpoint: core.StringPtr("$HOST/v2/spark/v3/instances/c7b3fccf-badb-46b0-b1ef-9b3154424021/spark_history_server"),
				SparkAccessEndpoint: core.StringPtr("$HOST/analytics-engine/details/spark-<instance_id>"),
				SparkJobsV4Endpoint: core.StringPtr("$HOST/v4/analytics_engines/c7b3fccf-badb-46b0-b1ef-9b3154424021/spark_applications"),
				SparkKernelEndpoint: core.StringPtr("$HOST/v4/analytics_engines/c7b3fccf-badb-46b0-b1ef-9b3154424021/jkg/api/kernels"),
				ViewHistoryServer: core.StringPtr("testString"),
				WxdApplicationEndpoint: core.StringPtr("$HOST/v1/1698311655308796/engines/spark817/applications"),
			}

			prestissimoEngineDetailsModel := &watsonxdatav2.PrestissimoEngineDetails{
				ApiKey: core.StringPtr("<api_key>"),
				ConnectionString: core.StringPtr("1.2.3.4"),
				Coordinator: prestissimoNodeDescriptionBodyModel,
				Endpoints: prestissimoEndpointsModel,
				InstanceID: core.StringPtr("instance_id"),
				ManagedBy: core.StringPtr("fully/self"),
				MetastoreHost: core.StringPtr("1.2.3.4"),
				SizeConfig: core.StringPtr("starter"),
				Worker: prestissimoNodeDescriptionBodyModel,
			}

			createPrestissimoEngineOptions := &watsonxdatav2.CreatePrestissimoEngineOptions{
				Origin: core.StringPtr("native"),
				AssociatedCatalogs: []string{"hive_data"},
				Description: core.StringPtr("prestissimo engine description"),
				EngineDetails: prestissimoEngineDetailsModel,
				EngineDisplayName: core.StringPtr("sampleEngine"),
				Region: core.StringPtr("us-south"),
				Tags: []string{"tag1", "tag2"},
				Version: core.StringPtr("1.2.3"),
				AuthInstanceID: core.StringPtr("testString"),
			}

			prestissimoEngine, response, err := watsonxDataService.CreatePrestissimoEngine(createPrestissimoEngineOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(201))
			Expect(prestissimoEngine).ToNot(BeNil())
		})
	})

	Describe(`GetPrestissimoEngine - Get prestissimo engine`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`GetPrestissimoEngine(getPrestissimoEngineOptions *GetPrestissimoEngineOptions)`, func() {
			getPrestissimoEngineOptions := &watsonxdatav2.GetPrestissimoEngineOptions{
				EngineID: core.StringPtr("testString"),
				AuthInstanceID: core.StringPtr("testString"),
			}

			prestissimoEngine, response, err := watsonxDataService.GetPrestissimoEngine(getPrestissimoEngineOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(prestissimoEngine).ToNot(BeNil())
		})
	})

	Describe(`UpdatePrestissimoEngine - Update prestissimo engine`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`UpdatePrestissimoEngine(updatePrestissimoEngineOptions *UpdatePrestissimoEngineOptions)`, func() {
			prestissimoEnginePropertiesCatalogModel := &watsonxdatav2.PrestissimoEnginePropertiesCatalog{
				CatalogName: []string{"testString"},
			}

			prestissimoNodeDescriptionBodyModel := &watsonxdatav2.PrestissimoNodeDescriptionBody{
				NodeType: core.StringPtr("worker"),
				Quantity: core.Int64Ptr(int64(38)),
			}

			enginePropertiesOaiGenConfigurationModel := &watsonxdatav2.EnginePropertiesOaiGenConfiguration{
				Coordinator: prestissimoNodeDescriptionBodyModel,
				Worker: prestissimoNodeDescriptionBodyModel,
			}

			prestissimoEnginePropertiesVeloxModel := &watsonxdatav2.PrestissimoEnginePropertiesVelox{
				VeloxProperty: []string{"testString"},
			}

			nodeDescriptionBodyModel := &watsonxdatav2.NodeDescriptionBody{
				NodeType: core.StringPtr("worker"),
				Quantity: core.Int64Ptr(int64(38)),
			}

			prestissimoEnginePropertiesOaiGen1JvmModel := &watsonxdatav2.PrestissimoEnginePropertiesOaiGen1Jvm{
				Coordinator: nodeDescriptionBodyModel,
			}

			prestissimoEngineEnginePropertiesModel := &watsonxdatav2.PrestissimoEngineEngineProperties{
				Catalog: prestissimoEnginePropertiesCatalogModel,
				Configuration: enginePropertiesOaiGenConfigurationModel,
				Velox: prestissimoEnginePropertiesVeloxModel,
				Jvm: prestissimoEnginePropertiesOaiGen1JvmModel,
			}

			removeEnginePropertiesConfigurationModel := &watsonxdatav2.RemoveEnginePropertiesConfiguration{
				Coordinator: []string{"testString"},
				Worker: []string{"testString"},
			}

			removeEnginePropertiesModel := &watsonxdatav2.RemoveEngineProperties{
				Catalog: prestissimoEnginePropertiesCatalogModel,
				Configuration: removeEnginePropertiesConfigurationModel,
				Jvm: removeEnginePropertiesConfigurationModel,
				Velox: []string{"testString"},
			}

			prestissimoEnginePatchModel := &watsonxdatav2.PrestissimoEnginePatch{
				Description: core.StringPtr("updated description for prestissimo engine"),
				EngineDisplayName: core.StringPtr("sampleEngine"),
				EngineProperties: prestissimoEngineEnginePropertiesModel,
				EngineRestart: core.StringPtr("force"),
				RemoveEngineProperties: removeEnginePropertiesModel,
				Tags: []string{"tag1", "tag2"},
			}
			prestissimoEnginePatchModelAsPatch, asPatchErr := prestissimoEnginePatchModel.AsPatch()
			Expect(asPatchErr).To(BeNil())

			updatePrestissimoEngineOptions := &watsonxdatav2.UpdatePrestissimoEngineOptions{
				EngineID: core.StringPtr("testString"),
				Body: prestissimoEnginePatchModelAsPatch,
				AuthInstanceID: core.StringPtr("testString"),
			}

			prestissimoEngine, response, err := watsonxDataService.UpdatePrestissimoEngine(updatePrestissimoEngineOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(prestissimoEngine).ToNot(BeNil())
		})
	})

	Describe(`ListPrestissimoEngineCatalogs - Get prestissimo engine catalogs`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`ListPrestissimoEngineCatalogs(listPrestissimoEngineCatalogsOptions *ListPrestissimoEngineCatalogsOptions)`, func() {
			listPrestissimoEngineCatalogsOptions := &watsonxdatav2.ListPrestissimoEngineCatalogsOptions{
				EngineID: core.StringPtr("testString"),
				AuthInstanceID: core.StringPtr("testString"),
			}

			catalogCollection, response, err := watsonxDataService.ListPrestissimoEngineCatalogs(listPrestissimoEngineCatalogsOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(catalogCollection).ToNot(BeNil())
		})
	})

	Describe(`AddPrestissimoEngineCatalogs - Associate catalogs to a prestissimo engine`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`AddPrestissimoEngineCatalogs(addPrestissimoEngineCatalogsOptions *AddPrestissimoEngineCatalogsOptions)`, func() {
			addPrestissimoEngineCatalogsOptions := &watsonxdatav2.AddPrestissimoEngineCatalogsOptions{
				EngineID: core.StringPtr("testString"),
				CatalogNames: core.StringPtr("testString"),
				AuthInstanceID: core.StringPtr("testString"),
			}

			catalogCollection, response, err := watsonxDataService.AddPrestissimoEngineCatalogs(addPrestissimoEngineCatalogsOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(201))
			Expect(catalogCollection).ToNot(BeNil())
		})
	})

	Describe(`GetPrestissimoEngineCatalog - Get prestissimo engine catalog`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`GetPrestissimoEngineCatalog(getPrestissimoEngineCatalogOptions *GetPrestissimoEngineCatalogOptions)`, func() {
			getPrestissimoEngineCatalogOptions := &watsonxdatav2.GetPrestissimoEngineCatalogOptions{
				EngineID: core.StringPtr("testString"),
				CatalogID: core.StringPtr("testString"),
				AuthInstanceID: core.StringPtr("testString"),
			}

			catalog, response, err := watsonxDataService.GetPrestissimoEngineCatalog(getPrestissimoEngineCatalogOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(catalog).ToNot(BeNil())
		})
	})

	Describe(`PausePrestissimoEngine - Pause prestissimo engine`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`PausePrestissimoEngine(pausePrestissimoEngineOptions *PausePrestissimoEngineOptions)`, func() {
			pausePrestissimoEngineOptions := &watsonxdatav2.PausePrestissimoEngineOptions{
				EngineID: core.StringPtr("testString"),
				AuthInstanceID: core.StringPtr("testString"),
			}

			successResponse, response, err := watsonxDataService.PausePrestissimoEngine(pausePrestissimoEngineOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(successResponse).ToNot(BeNil())
		})
	})

	Describe(`RunPrestissimoExplainStatement - Explain query`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`RunPrestissimoExplainStatement(runPrestissimoExplainStatementOptions *RunPrestissimoExplainStatementOptions)`, func() {
			runPrestissimoExplainStatementOptions := &watsonxdatav2.RunPrestissimoExplainStatementOptions{
				EngineID: core.StringPtr("testString"),
				Statement: core.StringPtr("show schemas in catalog_name"),
				Format: core.StringPtr("json"),
				Type: core.StringPtr("io"),
				AuthInstanceID: core.StringPtr("testString"),
			}

			resultPrestissimoExplainStatement, response, err := watsonxDataService.RunPrestissimoExplainStatement(runPrestissimoExplainStatementOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(resultPrestissimoExplainStatement).ToNot(BeNil())
		})
	})

	Describe(`RunPrestissimoExplainAnalyzeStatement - Explain analyze`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`RunPrestissimoExplainAnalyzeStatement(runPrestissimoExplainAnalyzeStatementOptions *RunPrestissimoExplainAnalyzeStatementOptions)`, func() {
			runPrestissimoExplainAnalyzeStatementOptions := &watsonxdatav2.RunPrestissimoExplainAnalyzeStatementOptions{
				EngineID: core.StringPtr("testString"),
				Statement: core.StringPtr("show schemas in catalog_name"),
				Verbose: core.BoolPtr(true),
				AuthInstanceID: core.StringPtr("testString"),
			}

			resultRunPrestissimoExplainAnalyzeStatement, response, err := watsonxDataService.RunPrestissimoExplainAnalyzeStatement(runPrestissimoExplainAnalyzeStatementOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(resultRunPrestissimoExplainAnalyzeStatement).ToNot(BeNil())
		})
	})

	Describe(`RestartPrestissimoEngine - Restart a prestissimo engine`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`RestartPrestissimoEngine(restartPrestissimoEngineOptions *RestartPrestissimoEngineOptions)`, func() {
			restartPrestissimoEngineOptions := &watsonxdatav2.RestartPrestissimoEngineOptions{
				EngineID: core.StringPtr("testString"),
				AuthInstanceID: core.StringPtr("testString"),
			}

			successResponse, response, err := watsonxDataService.RestartPrestissimoEngine(restartPrestissimoEngineOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(successResponse).ToNot(BeNil())
		})
	})

	Describe(`ResumePrestissimoEngine - Resume prestissimo engine`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`ResumePrestissimoEngine(resumePrestissimoEngineOptions *ResumePrestissimoEngineOptions)`, func() {
			resumePrestissimoEngineOptions := &watsonxdatav2.ResumePrestissimoEngineOptions{
				EngineID: core.StringPtr("testString"),
				AuthInstanceID: core.StringPtr("testString"),
			}

			successResponse, response, err := watsonxDataService.ResumePrestissimoEngine(resumePrestissimoEngineOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(successResponse).ToNot(BeNil())
		})
	})

	Describe(`ScalePrestissimoEngine - Scale a prestissimo engine`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`ScalePrestissimoEngine(scalePrestissimoEngineOptions *ScalePrestissimoEngineOptions)`, func() {
			prestissimoNodeDescriptionBodyModel := &watsonxdatav2.PrestissimoNodeDescriptionBody{
				NodeType: core.StringPtr("worker"),
				Quantity: core.Int64Ptr(int64(38)),
			}

			scalePrestissimoEngineOptions := &watsonxdatav2.ScalePrestissimoEngineOptions{
				EngineID: core.StringPtr("testString"),
				Coordinator: prestissimoNodeDescriptionBodyModel,
				Worker: prestissimoNodeDescriptionBodyModel,
				AuthInstanceID: core.StringPtr("testString"),
			}

			successResponse, response, err := watsonxDataService.ScalePrestissimoEngine(scalePrestissimoEngineOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(202))
			Expect(successResponse).ToNot(BeNil())
		})
	})

	Describe(`ListPrestoEngines - Get list of presto engines`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`ListPrestoEngines(listPrestoEnginesOptions *ListPrestoEnginesOptions)`, func() {
			listPrestoEnginesOptions := &watsonxdatav2.ListPrestoEnginesOptions{
				AuthInstanceID: core.StringPtr("testString"),
			}

			prestoEngineCollection, response, err := watsonxDataService.ListPrestoEngines(listPrestoEnginesOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(prestoEngineCollection).ToNot(BeNil())
		})
	})

	Describe(`CreatePrestoEngine - Create presto engine`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`CreatePrestoEngine(createPrestoEngineOptions *CreatePrestoEngineOptions)`, func() {
			nodeDescriptionBodyModel := &watsonxdatav2.NodeDescriptionBody{
				NodeType: core.StringPtr("worker"),
				Quantity: core.Int64Ptr(int64(38)),
			}

			engineDetailsBodyModel := &watsonxdatav2.EngineDetailsBody{
				ApiKey: core.StringPtr("<api_key>"),
				ConnectionString: core.StringPtr("1.2.3.4"),
				Coordinator: nodeDescriptionBodyModel,
				InstanceID: core.StringPtr("instance_id"),
				ManagedBy: core.StringPtr("fully/self"),
				SizeConfig: core.StringPtr("starter"),
				Worker: nodeDescriptionBodyModel,
			}

			createPrestoEngineOptions := &watsonxdatav2.CreatePrestoEngineOptions{
				Origin: core.StringPtr("native"),
				AssociatedCatalogs: []string{"iceberg_data", "hive_data"},
				Description: core.StringPtr("presto engine for running sql queries"),
				EngineDetails: engineDetailsBodyModel,
				EngineDisplayName: core.StringPtr("sampleEngine"),
				Region: core.StringPtr("us-south"),
				Tags: []string{"tag1", "tag2"},
				Version: core.StringPtr("1.2.3"),
				AuthInstanceID: core.StringPtr("testString"),
			}

			prestoEngine, response, err := watsonxDataService.CreatePrestoEngine(createPrestoEngineOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(201))
			Expect(prestoEngine).ToNot(BeNil())
		})
	})

	Describe(`GetPrestoEngine - Get presto engine`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`GetPrestoEngine(getPrestoEngineOptions *GetPrestoEngineOptions)`, func() {
			getPrestoEngineOptions := &watsonxdatav2.GetPrestoEngineOptions{
				EngineID: core.StringPtr("testString"),
				AuthInstanceID: core.StringPtr("testString"),
			}

			prestoEngine, response, err := watsonxDataService.GetPrestoEngine(getPrestoEngineOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(prestoEngine).ToNot(BeNil())
		})
	})

	Describe(`UpdatePrestoEngine - Update presto engine`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`UpdatePrestoEngine(updatePrestoEngineOptions *UpdatePrestoEngineOptions)`, func() {
			prestoEnginePropertiesCatalogModel := &watsonxdatav2.PrestoEnginePropertiesCatalog{
				CatalogName: core.StringPtr("testString"),
			}

			nodeDescriptionBodyModel := &watsonxdatav2.NodeDescriptionBody{
				NodeType: core.StringPtr("worker"),
				Quantity: core.Int64Ptr(int64(38)),
			}

			enginePropertiesOaiGen1ConfigurationModel := &watsonxdatav2.EnginePropertiesOaiGen1Configuration{
				Coordinator: nodeDescriptionBodyModel,
				Worker: nodeDescriptionBodyModel,
			}

			prestoEnginePropertiesGlobalModel := &watsonxdatav2.PrestoEnginePropertiesGlobal{
				GlobalProperty: core.StringPtr("enable-mixed-case-support:true"),
			}

			enginePropertiesOaiGen1JvmModel := &watsonxdatav2.EnginePropertiesOaiGen1Jvm{
				Coordinator: nodeDescriptionBodyModel,
				Worker: nodeDescriptionBodyModel,
			}

			prestoEngineEnginePropertiesModel := &watsonxdatav2.PrestoEngineEngineProperties{
				Catalog: prestoEnginePropertiesCatalogModel,
				Configuration: enginePropertiesOaiGen1ConfigurationModel,
				Global: prestoEnginePropertiesGlobalModel,
				Jvm: enginePropertiesOaiGen1JvmModel,
			}

			removeEnginePropertiesOaiGenConfigurationModel := &watsonxdatav2.RemoveEnginePropertiesOaiGenConfiguration{
				Coordinator: []string{"testString"},
				Worker: []string{"testString"},
			}

			removeEnginePropertiesOaiGenJvmModel := &watsonxdatav2.RemoveEnginePropertiesOaiGenJvm{
				Coordinator: []string{"testString"},
				Worker: []string{"testString"},
			}

			prestoEnginePatchRemoveEnginePropertiesModel := &watsonxdatav2.PrestoEnginePatchRemoveEngineProperties{
				Configuration: removeEnginePropertiesOaiGenConfigurationModel,
				Jvm: removeEnginePropertiesOaiGenJvmModel,
				Catalog: prestoEnginePropertiesCatalogModel,
			}

			prestoEnginePatchModel := &watsonxdatav2.PrestoEnginePatch{
				Description: core.StringPtr("updated description for presto engine"),
				EngineDisplayName: core.StringPtr("sampleEngine"),
				EngineProperties: prestoEngineEnginePropertiesModel,
				EngineRestart: core.StringPtr("force"),
				RemoveEngineProperties: prestoEnginePatchRemoveEnginePropertiesModel,
				Tags: []string{"tag1", "tag2"},
			}
			prestoEnginePatchModelAsPatch, asPatchErr := prestoEnginePatchModel.AsPatch()
			Expect(asPatchErr).To(BeNil())

			updatePrestoEngineOptions := &watsonxdatav2.UpdatePrestoEngineOptions{
				EngineID: core.StringPtr("testString"),
				Body: prestoEnginePatchModelAsPatch,
				AuthInstanceID: core.StringPtr("testString"),
			}

			prestoEngine, response, err := watsonxDataService.UpdatePrestoEngine(updatePrestoEngineOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(prestoEngine).ToNot(BeNil())
		})
	})

	Describe(`ListPrestoEngineCatalogs - Get presto engine catalogs`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`ListPrestoEngineCatalogs(listPrestoEngineCatalogsOptions *ListPrestoEngineCatalogsOptions)`, func() {
			listPrestoEngineCatalogsOptions := &watsonxdatav2.ListPrestoEngineCatalogsOptions{
				EngineID: core.StringPtr("testString"),
				AuthInstanceID: core.StringPtr("testString"),
			}

			catalogCollection, response, err := watsonxDataService.ListPrestoEngineCatalogs(listPrestoEngineCatalogsOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(catalogCollection).ToNot(BeNil())
		})
	})

	Describe(`AddPrestoEngineCatalogs - Associate catalogs to presto engine`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`AddPrestoEngineCatalogs(addPrestoEngineCatalogsOptions *AddPrestoEngineCatalogsOptions)`, func() {
			addPrestoEngineCatalogsOptions := &watsonxdatav2.AddPrestoEngineCatalogsOptions{
				EngineID: core.StringPtr("testString"),
				CatalogNames: core.StringPtr("testString"),
				AuthInstanceID: core.StringPtr("testString"),
			}

			catalogCollection, response, err := watsonxDataService.AddPrestoEngineCatalogs(addPrestoEngineCatalogsOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(201))
			Expect(catalogCollection).ToNot(BeNil())
		})
	})

	Describe(`GetPrestoEngineCatalog - Get presto engine catalog`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`GetPrestoEngineCatalog(getPrestoEngineCatalogOptions *GetPrestoEngineCatalogOptions)`, func() {
			getPrestoEngineCatalogOptions := &watsonxdatav2.GetPrestoEngineCatalogOptions{
				EngineID: core.StringPtr("testString"),
				CatalogID: core.StringPtr("testString"),
				AuthInstanceID: core.StringPtr("testString"),
			}

			catalog, response, err := watsonxDataService.GetPrestoEngineCatalog(getPrestoEngineCatalogOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(catalog).ToNot(BeNil())
		})
	})

	Describe(`PausePrestoEngine - Pause presto engine`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`PausePrestoEngine(pausePrestoEngineOptions *PausePrestoEngineOptions)`, func() {
			pausePrestoEngineOptions := &watsonxdatav2.PausePrestoEngineOptions{
				EngineID: core.StringPtr("testString"),
				AuthInstanceID: core.StringPtr("testString"),
			}

			createEnginePauseCreatedBody, response, err := watsonxDataService.PausePrestoEngine(pausePrestoEngineOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(createEnginePauseCreatedBody).ToNot(BeNil())
		})
	})

	Describe(`RunExplainStatement - Explain presto query`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`RunExplainStatement(runExplainStatementOptions *RunExplainStatementOptions)`, func() {
			runExplainStatementOptions := &watsonxdatav2.RunExplainStatementOptions{
				EngineID: core.StringPtr("testString"),
				Statement: core.StringPtr("show schemas in catalog_name"),
				Format: core.StringPtr("json"),
				Type: core.StringPtr("io"),
				AuthInstanceID: core.StringPtr("testString"),
			}

			runExplainStatementOkBody, response, err := watsonxDataService.RunExplainStatement(runExplainStatementOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(runExplainStatementOkBody).ToNot(BeNil())
		})
	})

	Describe(`RunExplainAnalyzeStatement - Explain presto analyze`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`RunExplainAnalyzeStatement(runExplainAnalyzeStatementOptions *RunExplainAnalyzeStatementOptions)`, func() {
			runExplainAnalyzeStatementOptions := &watsonxdatav2.RunExplainAnalyzeStatementOptions{
				EngineID: core.StringPtr("testString"),
				Statement: core.StringPtr("show schemas in catalog_name"),
				Verbose: core.BoolPtr(true),
				AuthInstanceID: core.StringPtr("testString"),
			}

			runExplainAnalyzeStatementOkBody, response, err := watsonxDataService.RunExplainAnalyzeStatement(runExplainAnalyzeStatementOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(runExplainAnalyzeStatementOkBody).ToNot(BeNil())
		})
	})

	Describe(`RestartPrestoEngine - Restart a presto engine`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`RestartPrestoEngine(restartPrestoEngineOptions *RestartPrestoEngineOptions)`, func() {
			restartPrestoEngineOptions := &watsonxdatav2.RestartPrestoEngineOptions{
				EngineID: core.StringPtr("testString"),
				AuthInstanceID: core.StringPtr("testString"),
			}

			createEngineRestartCreatedBody, response, err := watsonxDataService.RestartPrestoEngine(restartPrestoEngineOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(createEngineRestartCreatedBody).ToNot(BeNil())
		})
	})

	Describe(`ResumePrestoEngine - Resume presto engine`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`ResumePrestoEngine(resumePrestoEngineOptions *ResumePrestoEngineOptions)`, func() {
			resumePrestoEngineOptions := &watsonxdatav2.ResumePrestoEngineOptions{
				EngineID: core.StringPtr("testString"),
				AuthInstanceID: core.StringPtr("testString"),
			}

			createEngineResumeCreatedBody, response, err := watsonxDataService.ResumePrestoEngine(resumePrestoEngineOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(createEngineResumeCreatedBody).ToNot(BeNil())
		})
	})

	Describe(`ScalePrestoEngine - Scale a presto engine`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`ScalePrestoEngine(scalePrestoEngineOptions *ScalePrestoEngineOptions)`, func() {
			nodeDescriptionModel := &watsonxdatav2.NodeDescription{
				NodeType: core.StringPtr("worker"),
				Quantity: core.Int64Ptr(int64(38)),
			}

			scalePrestoEngineOptions := &watsonxdatav2.ScalePrestoEngineOptions{
				EngineID: core.StringPtr("testString"),
				Coordinator: nodeDescriptionModel,
				Worker: nodeDescriptionModel,
				AuthInstanceID: core.StringPtr("testString"),
			}

			createEngineScaleCreatedBody, response, err := watsonxDataService.ScalePrestoEngine(scalePrestoEngineOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(202))
			Expect(createEngineScaleCreatedBody).ToNot(BeNil())
		})
	})

	Describe(`ListSparkEngines - List all spark engines`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`ListSparkEngines(listSparkEnginesOptions *ListSparkEnginesOptions)`, func() {
			listSparkEnginesOptions := &watsonxdatav2.ListSparkEnginesOptions{
				AuthInstanceID: core.StringPtr("testString"),
			}

			sparkEngineCollection, response, err := watsonxDataService.ListSparkEngines(listSparkEnginesOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(sparkEngineCollection).ToNot(BeNil())
		})
	})

	Describe(`CreateSparkEngine - Create spark engine`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`CreateSparkEngine(createSparkEngineOptions *CreateSparkEngineOptions)`, func() {
			sparkDefaultConfigModel := &watsonxdatav2.SparkDefaultConfig{
				Config1: core.StringPtr("testString"),
				Config2: core.StringPtr("testString"),
			}

			sparkScaleConfigModel := &watsonxdatav2.SparkScaleConfig{
				AutoScaleEnabled: core.BoolPtr(true),
				CurrentNumberOfNodes: core.Int64Ptr(int64(2)),
				MaximumNumberOfNodes: core.Int64Ptr(int64(5)),
				MinimumNumberOfNodes: core.Int64Ptr(int64(1)),
				NodeType: core.StringPtr("small"),
				NumberOfNodes: core.Int64Ptr(int64(5)),
			}

			sparkEngineDetailsPrototypeModel := &watsonxdatav2.SparkEngineDetailsPrototype{
				ApiKey: core.StringPtr("apikey"),
				ConnectionString: core.StringPtr("1.2.3.4"),
				DefaultConfig: sparkDefaultConfigModel,
				DefaultVersion: core.StringPtr("3.3"),
				EngineHomeBucketDisplayName: core.StringPtr("test-spark-bucket"),
				EngineHomeBucketName: core.StringPtr("4fec0f8b-888a-4c16-8f38-250c8499e6ce-customer"),
				EngineHomePath: core.StringPtr("spark/spark1234"),
				EngineHomeVolumeID: core.StringPtr("1704979825978585"),
				EngineHomeVolumeName: core.StringPtr("my-volume"),
				EngineHomeVolumeStorageClass: core.StringPtr("nfs-client"),
				EngineHomeVolumeStorageSize: core.StringPtr("5Gi"),
				InstanceID: core.StringPtr("spark-id"),
				ManagedBy: core.StringPtr("fully/self"),
				ScaleConfig: sparkScaleConfigModel,
			}

			createSparkEngineOptions := &watsonxdatav2.CreateSparkEngineOptions{
				Origin: core.StringPtr("native"),
				AssociatedCatalogs: []string{"iceberg_data"},
				Description: core.StringPtr("testString"),
				EngineDetails: sparkEngineDetailsPrototypeModel,
				EngineDisplayName: core.StringPtr("test-native"),
				Status: core.StringPtr("testString"),
				Tags: []string{"testString"},
				AuthInstanceID: core.StringPtr("testString"),
			}

			sparkEngine, response, err := watsonxDataService.CreateSparkEngine(createSparkEngineOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(202))
			Expect(sparkEngine).ToNot(BeNil())
		})
	})

	Describe(`GetSparkEngine - Get spark engine`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`GetSparkEngine(getSparkEngineOptions *GetSparkEngineOptions)`, func() {
			getSparkEngineOptions := &watsonxdatav2.GetSparkEngineOptions{
				EngineID: core.StringPtr("testString"),
				AuthInstanceID: core.StringPtr("testString"),
			}

			sparkEngine, response, err := watsonxDataService.GetSparkEngine(getSparkEngineOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(sparkEngine).ToNot(BeNil())
		})
	})

	Describe(`UpdateSparkEngine - Update spark engine`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`UpdateSparkEngine(updateSparkEngineOptions *UpdateSparkEngineOptions)`, func() {
			updateSparkEngineBodyEngineDetailsModel := &watsonxdatav2.UpdateSparkEngineBodyEngineDetails{
				DefaultConfig: map[string]string{"key1": "testString"},
				DefaultVersion: core.StringPtr("3.4"),
			}

			updateSparkEngineBodyModel := &watsonxdatav2.UpdateSparkEngineBody{
				Description: core.StringPtr("Updated Description"),
				EngineDetails: updateSparkEngineBodyEngineDetailsModel,
				EngineDisplayName: core.StringPtr("Updated Display Name"),
				Tags: []string{"tag1", "tag2"},
			}
			updateSparkEngineBodyModelAsPatch, asPatchErr := updateSparkEngineBodyModel.AsPatch()
			Expect(asPatchErr).To(BeNil())

			updateSparkEngineOptions := &watsonxdatav2.UpdateSparkEngineOptions{
				EngineID: core.StringPtr("testString"),
				Body: updateSparkEngineBodyModelAsPatch,
				AuthInstanceID: core.StringPtr("testString"),
			}

			sparkEngine, response, err := watsonxDataService.UpdateSparkEngine(updateSparkEngineOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(sparkEngine).ToNot(BeNil())
		})
	})

	Describe(`ListSparkEngineApplications - List all applications in a spark engine`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`ListSparkEngineApplications(listSparkEngineApplicationsOptions *ListSparkEngineApplicationsOptions)`, func() {
			listSparkEngineApplicationsOptions := &watsonxdatav2.ListSparkEngineApplicationsOptions{
				EngineID: core.StringPtr("testString"),
				AuthInstanceID: core.StringPtr("testString"),
				State: []string{"testString"},
			}

			sparkEngineApplicationStatusCollection, response, err := watsonxDataService.ListSparkEngineApplications(listSparkEngineApplicationsOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(sparkEngineApplicationStatusCollection).ToNot(BeNil())
		})
	})

	Describe(`CreateSparkEngineApplication - Submit engine applications`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`CreateSparkEngineApplication(createSparkEngineApplicationOptions *CreateSparkEngineApplicationOptions)`, func() {
			sparkApplicationConfigModel := &watsonxdatav2.SparkApplicationConfig{
				SparkSampleConfigProperpty: core.StringPtr("testString"),
			}

			sparkApplicationEnvModel := &watsonxdatav2.SparkApplicationEnv{
				SampleEnvKey: core.StringPtr("testString"),
			}

			sparkApplicationDetailsModel := &watsonxdatav2.SparkApplicationDetails{
				Application: core.StringPtr("/opt/ibm/spark/examples/src/main/python/wordcount.py"),
				Arguments: []string{"/opt/ibm/spark/examples/src/main/resources/people.txt"},
				Class: core.StringPtr("org.apache.spark.examples.SparkPi"),
				Conf: sparkApplicationConfigModel,
				Env: sparkApplicationEnvModel,
				Files: core.StringPtr("s3://mybucket/myfile.txt"),
				Jars: core.StringPtr("testString"),
				Name: core.StringPtr("SparkApplicaton1"),
				Packages: core.StringPtr("org.apache.spark:example_1.2.3"),
				Repositories: core.StringPtr("https://repo1.maven.org/maven2/"),
				SparkVersion: core.StringPtr("3.3"),
			}

			sparkVolumeDetailsModel := &watsonxdatav2.SparkVolumeDetails{
				MountPath: core.StringPtr("/mount/path"),
				Name: core.StringPtr("my-volume"),
				ReadOnly: core.BoolPtr(true),
				SourceSubPath: core.StringPtr("/source/path"),
			}

			createSparkEngineApplicationOptions := &watsonxdatav2.CreateSparkEngineApplicationOptions{
				EngineID: core.StringPtr("testString"),
				ApplicationDetails: sparkApplicationDetailsModel,
				JobEndpoint: core.StringPtr("testString"),
				ServiceInstanceID: core.StringPtr("testString"),
				Type: core.StringPtr("iae"),
				Volumes: []watsonxdatav2.SparkVolumeDetails{*sparkVolumeDetailsModel},
				AuthInstanceID: core.StringPtr("testString"),
				State: []string{"testString"},
			}

			sparkEngineApplicationStatus, response, err := watsonxDataService.CreateSparkEngineApplication(createSparkEngineApplicationOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(201))
			Expect(sparkEngineApplicationStatus).ToNot(BeNil())
		})
	})

	Describe(`GetSparkEngineApplicationStatus - Get spark application`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`GetSparkEngineApplicationStatus(getSparkEngineApplicationStatusOptions *GetSparkEngineApplicationStatusOptions)`, func() {
			getSparkEngineApplicationStatusOptions := &watsonxdatav2.GetSparkEngineApplicationStatusOptions{
				EngineID: core.StringPtr("testString"),
				ApplicationID: core.StringPtr("testString"),
				AuthInstanceID: core.StringPtr("testString"),
			}

			sparkEngineApplicationStatus, response, err := watsonxDataService.GetSparkEngineApplicationStatus(getSparkEngineApplicationStatusOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(sparkEngineApplicationStatus).ToNot(BeNil())
		})
	})

	Describe(`ListSparkEngineCatalogs - Get spark engine catalogs`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`ListSparkEngineCatalogs(listSparkEngineCatalogsOptions *ListSparkEngineCatalogsOptions)`, func() {
			listSparkEngineCatalogsOptions := &watsonxdatav2.ListSparkEngineCatalogsOptions{
				EngineID: core.StringPtr("testString"),
				AuthInstanceID: core.StringPtr("testString"),
			}

			catalogCollection, response, err := watsonxDataService.ListSparkEngineCatalogs(listSparkEngineCatalogsOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(catalogCollection).ToNot(BeNil())
		})
	})

	Describe(`AddSparkEngineCatalogs - Associate catalogs to spark engine`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`AddSparkEngineCatalogs(addSparkEngineCatalogsOptions *AddSparkEngineCatalogsOptions)`, func() {
			addSparkEngineCatalogsOptions := &watsonxdatav2.AddSparkEngineCatalogsOptions{
				EngineID: core.StringPtr("testString"),
				CatalogNames: core.StringPtr("testString"),
				AuthInstanceID: core.StringPtr("testString"),
			}

			catalogCollection, response, err := watsonxDataService.AddSparkEngineCatalogs(addSparkEngineCatalogsOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(201))
			Expect(catalogCollection).ToNot(BeNil())
		})
	})

	Describe(`GetSparkEngineCatalog - Get spark engine catalog`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`GetSparkEngineCatalog(getSparkEngineCatalogOptions *GetSparkEngineCatalogOptions)`, func() {
			getSparkEngineCatalogOptions := &watsonxdatav2.GetSparkEngineCatalogOptions{
				EngineID: core.StringPtr("testString"),
				CatalogID: core.StringPtr("testString"),
				AuthInstanceID: core.StringPtr("testString"),
			}

			catalog, response, err := watsonxDataService.GetSparkEngineCatalog(getSparkEngineCatalogOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(catalog).ToNot(BeNil())
		})
	})

	Describe(`GetSparkEngineHistoryServer - Get spark history server`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`GetSparkEngineHistoryServer(getSparkEngineHistoryServerOptions *GetSparkEngineHistoryServerOptions)`, func() {
			getSparkEngineHistoryServerOptions := &watsonxdatav2.GetSparkEngineHistoryServerOptions{
				EngineID: core.StringPtr("testString"),
				AuthInstanceID: core.StringPtr("testString"),
			}

			sparkHistoryServer, response, err := watsonxDataService.GetSparkEngineHistoryServer(getSparkEngineHistoryServerOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(sparkHistoryServer).ToNot(BeNil())
		})
	})

	Describe(`StartSparkEngineHistoryServer - Start spark history server`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`StartSparkEngineHistoryServer(startSparkEngineHistoryServerOptions *StartSparkEngineHistoryServerOptions)`, func() {
			startSparkEngineHistoryServerOptions := &watsonxdatav2.StartSparkEngineHistoryServerOptions{
				EngineID: core.StringPtr("testString"),
				Cores: core.StringPtr("1"),
				Memory: core.StringPtr("4G"),
				AuthInstanceID: core.StringPtr("testString"),
			}

			sparkHistoryServer, response, err := watsonxDataService.StartSparkEngineHistoryServer(startSparkEngineHistoryServerOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(201))
			Expect(sparkHistoryServer).ToNot(BeNil())
		})
	})

	Describe(`CreateSparkEnginePause - Pause engine`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`CreateSparkEnginePause(createSparkEnginePauseOptions *CreateSparkEnginePauseOptions)`, func() {
			createSparkEnginePauseOptions := &watsonxdatav2.CreateSparkEnginePauseOptions{
				EngineID: core.StringPtr("testString"),
				AuthInstanceID: core.StringPtr("testString"),
			}

			successResponse, response, err := watsonxDataService.CreateSparkEnginePause(createSparkEnginePauseOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(201))
			Expect(successResponse).ToNot(BeNil())
		})
	})

	Describe(`CreateSparkEngineResume - Resume engine`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`CreateSparkEngineResume(createSparkEngineResumeOptions *CreateSparkEngineResumeOptions)`, func() {
			createSparkEngineResumeOptions := &watsonxdatav2.CreateSparkEngineResumeOptions{
				EngineID: core.StringPtr("testString"),
				AuthInstanceID: core.StringPtr("testString"),
			}

			successResponse, response, err := watsonxDataService.CreateSparkEngineResume(createSparkEngineResumeOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(201))
			Expect(successResponse).ToNot(BeNil())
		})
	})

	Describe(`CreateSparkEngineScale - Scale Spark engine`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`CreateSparkEngineScale(createSparkEngineScaleOptions *CreateSparkEngineScaleOptions)`, func() {
			createSparkEngineScaleOptions := &watsonxdatav2.CreateSparkEngineScaleOptions{
				EngineID: core.StringPtr("testString"),
				NumberOfNodes: core.Int64Ptr(int64(2)),
				AuthInstanceID: core.StringPtr("testString"),
			}

			successResponse, response, err := watsonxDataService.CreateSparkEngineScale(createSparkEngineScaleOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(201))
			Expect(successResponse).ToNot(BeNil())
		})
	})

	Describe(`ListSparkVersions - List spark version`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`ListSparkVersions(listSparkVersionsOptions *ListSparkVersionsOptions)`, func() {
			listSparkVersionsOptions := &watsonxdatav2.ListSparkVersionsOptions{
				AuthInstanceID: core.StringPtr("testString"),
			}

			listSparkVersionsOkBody, response, err := watsonxDataService.ListSparkVersions(listSparkVersionsOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(listSparkVersionsOkBody).ToNot(BeNil())
		})
	})

	Describe(`ListCatalogs - List all registered catalogs`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`ListCatalogs(listCatalogsOptions *ListCatalogsOptions)`, func() {
			listCatalogsOptions := &watsonxdatav2.ListCatalogsOptions{
				AuthInstanceID: core.StringPtr("testString"),
			}

			catalogCollection, response, err := watsonxDataService.ListCatalogs(listCatalogsOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(catalogCollection).ToNot(BeNil())
		})
	})

	Describe(`GetCatalog - Get catalog properties by catalog_id`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`GetCatalog(getCatalogOptions *GetCatalogOptions)`, func() {
			getCatalogOptions := &watsonxdatav2.GetCatalogOptions{
				CatalogID: core.StringPtr("testString"),
				AuthInstanceID: core.StringPtr("testString"),
			}

			catalog, response, err := watsonxDataService.GetCatalog(getCatalogOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(catalog).ToNot(BeNil())
		})
	})

	Describe(`ListSchemas - List all schemas`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`ListSchemas(listSchemasOptions *ListSchemasOptions)`, func() {
			listSchemasOptions := &watsonxdatav2.ListSchemasOptions{
				EngineID: core.StringPtr("testString"),
				CatalogID: core.StringPtr("testString"),
				AuthInstanceID: core.StringPtr("testString"),
			}

			listSchemasOkBody, response, err := watsonxDataService.ListSchemas(listSchemasOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(listSchemasOkBody).ToNot(BeNil())
		})
	})

	Describe(`CreateSchema - Create schema`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`CreateSchema(createSchemaOptions *CreateSchemaOptions)`, func() {
			createSchemaOptions := &watsonxdatav2.CreateSchemaOptions{
				EngineID: core.StringPtr("testString"),
				CatalogID: core.StringPtr("testString"),
				CustomPath: core.StringPtr("sample-path"),
				SchemaName: core.StringPtr("SampleSchema1"),
				BucketName: core.StringPtr("sample-bucket"),
				AuthInstanceID: core.StringPtr("testString"),
			}

			createSchemaCreatedBody, response, err := watsonxDataService.CreateSchema(createSchemaOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(201))
			Expect(createSchemaCreatedBody).ToNot(BeNil())
		})
	})

	Describe(`ListTables - List all tables`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`ListTables(listTablesOptions *ListTablesOptions)`, func() {
			listTablesOptions := &watsonxdatav2.ListTablesOptions{
				CatalogID: core.StringPtr("testString"),
				SchemaID: core.StringPtr("testString"),
				EngineID: core.StringPtr("testString"),
				AuthInstanceID: core.StringPtr("testString"),
			}

			tableCollection, response, err := watsonxDataService.ListTables(listTablesOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(tableCollection).ToNot(BeNil())
		})
	})

	Describe(`GetTable - Get table details`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`GetTable(getTableOptions *GetTableOptions)`, func() {
			getTableOptions := &watsonxdatav2.GetTableOptions{
				CatalogID: core.StringPtr("testString"),
				SchemaID: core.StringPtr("testString"),
				TableID: core.StringPtr("testString"),
				EngineID: core.StringPtr("testString"),
				AuthInstanceID: core.StringPtr("testString"),
			}

			table, response, err := watsonxDataService.GetTable(getTableOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(table).ToNot(BeNil())
		})
	})

	Describe(`RenameTable - Rename table`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`RenameTable(renameTableOptions *RenameTableOptions)`, func() {
			tablePatchModel := &watsonxdatav2.TablePatch{
				TableName: core.StringPtr("updated_table_name"),
			}
			tablePatchModelAsPatch, asPatchErr := tablePatchModel.AsPatch()
			Expect(asPatchErr).To(BeNil())

			renameTableOptions := &watsonxdatav2.RenameTableOptions{
				CatalogID: core.StringPtr("testString"),
				SchemaID: core.StringPtr("testString"),
				TableID: core.StringPtr("testString"),
				EngineID: core.StringPtr("testString"),
				Body: tablePatchModelAsPatch,
				AuthInstanceID: core.StringPtr("testString"),
			}

			table, response, err := watsonxDataService.RenameTable(renameTableOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(table).ToNot(BeNil())
		})
	})

	Describe(`ListColumns - List all columns of a table`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`ListColumns(listColumnsOptions *ListColumnsOptions)`, func() {
			listColumnsOptions := &watsonxdatav2.ListColumnsOptions{
				EngineID: core.StringPtr("testString"),
				CatalogID: core.StringPtr("testString"),
				SchemaID: core.StringPtr("testString"),
				TableID: core.StringPtr("testString"),
				AuthInstanceID: core.StringPtr("testString"),
			}

			columnCollection, response, err := watsonxDataService.ListColumns(listColumnsOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(columnCollection).ToNot(BeNil())
		})
	})

	Describe(`CreateColumns - Add column(s)`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`CreateColumns(createColumnsOptions *CreateColumnsOptions)`, func() {
			columnModel := &watsonxdatav2.Column{
				ColumnName: core.StringPtr("expenses"),
				Comment: core.StringPtr("expenses column"),
				Extra: core.StringPtr("varchar"),
				Length: core.StringPtr("30"),
				Scale: core.StringPtr("2"),
				Type: core.StringPtr("varchar"),
			}

			createColumnsOptions := &watsonxdatav2.CreateColumnsOptions{
				EngineID: core.StringPtr("testString"),
				CatalogID: core.StringPtr("testString"),
				SchemaID: core.StringPtr("testString"),
				TableID: core.StringPtr("testString"),
				Columns: []watsonxdatav2.Column{*columnModel},
				AuthInstanceID: core.StringPtr("testString"),
			}

			columnCollection, response, err := watsonxDataService.CreateColumns(createColumnsOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(201))
			Expect(columnCollection).ToNot(BeNil())
		})
	})

	Describe(`UpdateColumn - Alter column`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`UpdateColumn(updateColumnOptions *UpdateColumnOptions)`, func() {
			columnPatchModel := &watsonxdatav2.ColumnPatch{
				ColumnName: core.StringPtr("expenses"),
			}
			columnPatchModelAsPatch, asPatchErr := columnPatchModel.AsPatch()
			Expect(asPatchErr).To(BeNil())

			updateColumnOptions := &watsonxdatav2.UpdateColumnOptions{
				EngineID: core.StringPtr("testString"),
				CatalogID: core.StringPtr("testString"),
				SchemaID: core.StringPtr("testString"),
				TableID: core.StringPtr("testString"),
				ColumnID: core.StringPtr("testString"),
				Body: columnPatchModelAsPatch,
				AuthInstanceID: core.StringPtr("testString"),
			}

			column, response, err := watsonxDataService.UpdateColumn(updateColumnOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(column).ToNot(BeNil())
		})
	})

	Describe(`ListTableSnapshots - Get table snapshots`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`ListTableSnapshots(listTableSnapshotsOptions *ListTableSnapshotsOptions)`, func() {
			listTableSnapshotsOptions := &watsonxdatav2.ListTableSnapshotsOptions{
				EngineID: core.StringPtr("testString"),
				CatalogID: core.StringPtr("testString"),
				SchemaID: core.StringPtr("testString"),
				TableID: core.StringPtr("testString"),
				AuthInstanceID: core.StringPtr("testString"),
			}

			tableSnapshotCollection, response, err := watsonxDataService.ListTableSnapshots(listTableSnapshotsOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(tableSnapshotCollection).ToNot(BeNil())
		})
	})

	Describe(`RollbackTable - Rollback table to snapshot`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`RollbackTable(rollbackTableOptions *RollbackTableOptions)`, func() {
			rollbackTableOptions := &watsonxdatav2.RollbackTableOptions{
				EngineID: core.StringPtr("testString"),
				CatalogID: core.StringPtr("testString"),
				SchemaID: core.StringPtr("testString"),
				TableID: core.StringPtr("testString"),
				SnapshotID: core.StringPtr("testString"),
				AuthInstanceID: core.StringPtr("testString"),
			}

			replaceSnapshotCreatedBody, response, err := watsonxDataService.RollbackTable(rollbackTableOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(201))
			Expect(replaceSnapshotCreatedBody).ToNot(BeNil())
		})
	})

	Describe(`UpdateSyncCatalog - External Iceberg table registration`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`UpdateSyncCatalog(updateSyncCatalogOptions *UpdateSyncCatalogOptions)`, func() {
			syncCatalogsModel := &watsonxdatav2.SyncCatalogs{
				AutoAddNewTables: core.BoolPtr(true),
				SyncIcebergMd: core.BoolPtr(true),
			}
			syncCatalogsModelAsPatch, asPatchErr := syncCatalogsModel.AsPatch()
			Expect(asPatchErr).To(BeNil())

			updateSyncCatalogOptions := &watsonxdatav2.UpdateSyncCatalogOptions{
				CatalogID: core.StringPtr("testString"),
				Body: syncCatalogsModelAsPatch,
				AuthInstanceID: core.StringPtr("testString"),
			}

			updateSyncCatalogOkBody, response, err := watsonxDataService.UpdateSyncCatalog(updateSyncCatalogOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(updateSyncCatalogOkBody).ToNot(BeNil())
		})
	})

	Describe(`ListMilvusServices - Get list of milvus services`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`ListMilvusServices(listMilvusServicesOptions *ListMilvusServicesOptions)`, func() {
			listMilvusServicesOptions := &watsonxdatav2.ListMilvusServicesOptions{
				AuthInstanceID: core.StringPtr("testString"),
			}

			milvusServiceCollection, response, err := watsonxDataService.ListMilvusServices(listMilvusServicesOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(milvusServiceCollection).ToNot(BeNil())
		})
	})

	Describe(`CreateMilvusService - Create milvus service`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`CreateMilvusService(createMilvusServiceOptions *CreateMilvusServiceOptions)`, func() {
			createMilvusServiceOptions := &watsonxdatav2.CreateMilvusServiceOptions{
				Origin: core.StringPtr("native"),
				Description: core.StringPtr("milvus service for running sql queries"),
				ServiceDisplayName: core.StringPtr("sampleService"),
				Tags: []string{"tag1", "tag2"},
				AuthInstanceID: core.StringPtr("testString"),
			}

			milvusService, response, err := watsonxDataService.CreateMilvusService(createMilvusServiceOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(201))
			Expect(milvusService).ToNot(BeNil())
		})
	})

	Describe(`GetMilvusService - Get milvus service`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`GetMilvusService(getMilvusServiceOptions *GetMilvusServiceOptions)`, func() {
			getMilvusServiceOptions := &watsonxdatav2.GetMilvusServiceOptions{
				ServiceID: core.StringPtr("testString"),
				AuthInstanceID: core.StringPtr("testString"),
			}

			milvusService, response, err := watsonxDataService.GetMilvusService(getMilvusServiceOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(milvusService).ToNot(BeNil())
		})
	})

	Describe(`UpdateMilvusService - Update milvus service`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`UpdateMilvusService(updateMilvusServiceOptions *UpdateMilvusServiceOptions)`, func() {
			milvusServicePatchModel := &watsonxdatav2.MilvusServicePatch{
				Description: core.StringPtr("updated description for milvus service"),
				ServiceDisplayName: core.StringPtr("sampleService"),
				Tags: []string{"tag1", "tag2"},
			}
			milvusServicePatchModelAsPatch, asPatchErr := milvusServicePatchModel.AsPatch()
			Expect(asPatchErr).To(BeNil())

			updateMilvusServiceOptions := &watsonxdatav2.UpdateMilvusServiceOptions{
				ServiceID: core.StringPtr("testString"),
				Body: milvusServicePatchModelAsPatch,
				AuthInstanceID: core.StringPtr("testString"),
			}

			milvusService, response, err := watsonxDataService.UpdateMilvusService(updateMilvusServiceOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(milvusService).ToNot(BeNil())
		})
	})

	Describe(`ListIngestionJobs - Get ingestion jobs`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`ListIngestionJobs(listIngestionJobsOptions *ListIngestionJobsOptions)`, func() {
			listIngestionJobsOptions := &watsonxdatav2.ListIngestionJobsOptions{
				AuthInstanceID: core.StringPtr("testString"),
				Page: core.Int64Ptr(int64(1)),
				JobsPerPage: core.Int64Ptr(int64(1)),
			}

			ingestionJobCollection, response, err := watsonxDataService.ListIngestionJobs(listIngestionJobsOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(ingestionJobCollection).ToNot(BeNil())
		})
	})

	Describe(`CreateIngestionJobs - Create an ingestion job`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`CreateIngestionJobs(createIngestionJobsOptions *CreateIngestionJobsOptions)`, func() {
			ingestionJobPrototypeCsvPropertyModel := &watsonxdatav2.IngestionJobPrototypeCsvProperty{
				Encoding: core.StringPtr("utf-8"),
				EscapeCharacter: core.StringPtr("\\\\"),
				FieldDelimiter: core.StringPtr(","),
				Header: core.BoolPtr(true),
				LineDelimiter: core.StringPtr("\\n"),
			}

			ingestionJobPrototypeExecuteConfigModel := &watsonxdatav2.IngestionJobPrototypeExecuteConfig{
				DriverCores: core.Int64Ptr(int64(1)),
				DriverMemory: core.StringPtr("2G"),
				ExecutorCores: core.Int64Ptr(int64(1)),
				ExecutorMemory: core.StringPtr("2G"),
				NumExecutors: core.Int64Ptr(int64(1)),
			}

			createIngestionJobsOptions := &watsonxdatav2.CreateIngestionJobsOptions{
				AuthInstanceID: core.StringPtr("testString"),
				JobID: core.StringPtr("ingestion-1699459946935"),
				SourceDataFiles: core.StringPtr("s3://demobucket/data/yellow_tripdata_2022-01.parquet"),
				TargetTable: core.StringPtr("demodb.test.targettable"),
				Username: core.StringPtr("user1"),
				CreateIfNotExist: core.BoolPtr(false),
				CsvProperty: ingestionJobPrototypeCsvPropertyModel,
				EngineID: core.StringPtr("spark123"),
				ExecuteConfig: ingestionJobPrototypeExecuteConfigModel,
				PartitionBy: core.StringPtr("col1, col2"),
				Schema: core.StringPtr("{\"type\":\"struct\",\"schema-id\":0,\"fields\":[{\"id\":1,\"name\":\"ID\",\"required\":true,\"type\":\"int\"},{\"id\":2,\"name\":\"Name\",\"required\":true,\"type\":\"string\"}]}"),
				SourceFileType: core.StringPtr("csv"),
				ValidateCsvHeader: core.BoolPtr(false),
			}

			ingestionJob, response, err := watsonxDataService.CreateIngestionJobs(createIngestionJobsOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(202))
			Expect(ingestionJob).ToNot(BeNil())
		})
	})

	Describe(`CreateIngestionJobsLocalFiles - Create an ingestion job for user local files`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`CreateIngestionJobsLocalFiles(createIngestionJobsLocalFilesOptions *CreateIngestionJobsLocalFilesOptions)`, func() {
			createIngestionJobsLocalFilesOptions := &watsonxdatav2.CreateIngestionJobsLocalFilesOptions{
				AuthInstanceID: core.StringPtr("testString"),
				SourceDataFile: CreateMockReader("This is a mock file."),
				TargetTable: core.StringPtr("testString"),
				JobID: core.StringPtr("testString"),
				Username: core.StringPtr("testString"),
				SourceDataFileContentType: core.StringPtr("testString"),
				SourceFileType: core.StringPtr("csv"),
				CsvProperty: core.StringPtr("testString"),
				CreateIfNotExist: core.BoolPtr(false),
				ValidateCsvHeader: core.BoolPtr(false),
				ExecuteConfig: core.StringPtr("testString"),
				EngineID: core.StringPtr("testString"),
			}

			ingestionJob, response, err := watsonxDataService.CreateIngestionJobsLocalFiles(createIngestionJobsLocalFilesOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(202))
			Expect(ingestionJob).ToNot(BeNil())
		})
	})

	Describe(`GetIngestionJob - Get ingestion job`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`GetIngestionJob(getIngestionJobOptions *GetIngestionJobOptions)`, func() {
			getIngestionJobOptions := &watsonxdatav2.GetIngestionJobOptions{
				JobID: core.StringPtr("testString"),
				AuthInstanceID: core.StringPtr("testString"),
			}

			ingestionJob, response, err := watsonxDataService.GetIngestionJob(getIngestionJobOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(ingestionJob).ToNot(BeNil())
		})
	})

	Describe(`CreatePreviewIngestionFile - Generate a preview of source file(s)`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`CreatePreviewIngestionFile(createPreviewIngestionFileOptions *CreatePreviewIngestionFileOptions)`, func() {
			previewIngestionFilePrototypeCsvPropertyModel := &watsonxdatav2.PreviewIngestionFilePrototypeCsvProperty{
				Encoding: core.StringPtr("utf-8"),
				EscapeCharacter: core.StringPtr("\\\\"),
				FieldDelimiter: core.StringPtr(","),
				Header: core.BoolPtr(true),
				LineDelimiter: core.StringPtr("\\n"),
			}

			createPreviewIngestionFileOptions := &watsonxdatav2.CreatePreviewIngestionFileOptions{
				AuthInstanceID: core.StringPtr("testString"),
				SourceDataFiles: core.StringPtr("s3://demobucket/data/yellow_tripdata_2022-01.parquet"),
				CsvProperty: previewIngestionFilePrototypeCsvPropertyModel,
				SourceFileType: core.StringPtr("csv"),
			}

			previewIngestionFile, response, err := watsonxDataService.CreatePreviewIngestionFile(createPreviewIngestionFileOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(201))
			Expect(previewIngestionFile).ToNot(BeNil())
		})
	})

	Describe(`DeregisterBucket - Deregister Bucket`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`DeregisterBucket(deregisterBucketOptions *DeregisterBucketOptions)`, func() {
			deregisterBucketOptions := &watsonxdatav2.DeregisterBucketOptions{
				BucketID: core.StringPtr("testString"),
				AuthInstanceID: core.StringPtr("testString"),
			}

			response, err := watsonxDataService.DeregisterBucket(deregisterBucketOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(204))
		})
	})

	Describe(`DeleteDeactivateBucket - Deactivate Bucket`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`DeleteDeactivateBucket(deleteDeactivateBucketOptions *DeleteDeactivateBucketOptions)`, func() {
			deleteDeactivateBucketOptions := &watsonxdatav2.DeleteDeactivateBucketOptions{
				BucketID: core.StringPtr("testString"),
				AuthInstanceID: core.StringPtr("testString"),
			}

			response, err := watsonxDataService.DeleteDeactivateBucket(deleteDeactivateBucketOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(204))
		})
	})

	Describe(`DeleteDatabaseCatalog - Delete database`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`DeleteDatabaseCatalog(deleteDatabaseCatalogOptions *DeleteDatabaseCatalogOptions)`, func() {
			deleteDatabaseCatalogOptions := &watsonxdatav2.DeleteDatabaseCatalogOptions{
				DatabaseID: core.StringPtr("testString"),
				AuthInstanceID: core.StringPtr("testString"),
			}

			response, err := watsonxDataService.DeleteDatabaseCatalog(deleteDatabaseCatalogOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(204))
		})
	})

	Describe(`DeleteOtherEngine - Delete engine`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`DeleteOtherEngine(deleteOtherEngineOptions *DeleteOtherEngineOptions)`, func() {
			deleteOtherEngineOptions := &watsonxdatav2.DeleteOtherEngineOptions{
				EngineID: core.StringPtr("testString"),
				AuthInstanceID: core.StringPtr("testString"),
			}

			response, err := watsonxDataService.DeleteOtherEngine(deleteOtherEngineOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(204))
		})
	})

	Describe(`DeleteDb2Engine - Delete db2 engine`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`DeleteDb2Engine(deleteDb2EngineOptions *DeleteDb2EngineOptions)`, func() {
			deleteDb2EngineOptions := &watsonxdatav2.DeleteDb2EngineOptions{
				EngineID: core.StringPtr("testString"),
				AuthInstanceID: core.StringPtr("testString"),
			}

			response, err := watsonxDataService.DeleteDb2Engine(deleteDb2EngineOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(204))
		})
	})

	Describe(`DeleteNetezzaEngine - Delete netezza engine`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`DeleteNetezzaEngine(deleteNetezzaEngineOptions *DeleteNetezzaEngineOptions)`, func() {
			deleteNetezzaEngineOptions := &watsonxdatav2.DeleteNetezzaEngineOptions{
				EngineID: core.StringPtr("testString"),
				AuthInstanceID: core.StringPtr("testString"),
			}

			response, err := watsonxDataService.DeleteNetezzaEngine(deleteNetezzaEngineOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(204))
		})
	})

	Describe(`DeletePrestissimoEngine - Delete prestissimo engine`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`DeletePrestissimoEngine(deletePrestissimoEngineOptions *DeletePrestissimoEngineOptions)`, func() {
			deletePrestissimoEngineOptions := &watsonxdatav2.DeletePrestissimoEngineOptions{
				EngineID: core.StringPtr("testString"),
				AuthInstanceID: core.StringPtr("testString"),
			}

			response, err := watsonxDataService.DeletePrestissimoEngine(deletePrestissimoEngineOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(204))
		})
	})

	Describe(`DeletePrestissimoEngineCatalogs - Disassociate catalogs from a prestissimo engine`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`DeletePrestissimoEngineCatalogs(deletePrestissimoEngineCatalogsOptions *DeletePrestissimoEngineCatalogsOptions)`, func() {
			deletePrestissimoEngineCatalogsOptions := &watsonxdatav2.DeletePrestissimoEngineCatalogsOptions{
				EngineID: core.StringPtr("testString"),
				CatalogNames: core.StringPtr("testString"),
				AuthInstanceID: core.StringPtr("testString"),
			}

			response, err := watsonxDataService.DeletePrestissimoEngineCatalogs(deletePrestissimoEngineCatalogsOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(204))
		})
	})

	Describe(`DeleteEngine - Delete presto engine`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`DeleteEngine(deleteEngineOptions *DeleteEngineOptions)`, func() {
			deleteEngineOptions := &watsonxdatav2.DeleteEngineOptions{
				EngineID: core.StringPtr("testString"),
				AuthInstanceID: core.StringPtr("testString"),
			}

			response, err := watsonxDataService.DeleteEngine(deleteEngineOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(204))
		})
	})

	Describe(`DeletePrestoEngineCatalogs - Disassociate catalogs from a presto engine`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`DeletePrestoEngineCatalogs(deletePrestoEngineCatalogsOptions *DeletePrestoEngineCatalogsOptions)`, func() {
			deletePrestoEngineCatalogsOptions := &watsonxdatav2.DeletePrestoEngineCatalogsOptions{
				EngineID: core.StringPtr("testString"),
				CatalogNames: core.StringPtr("testString"),
				AuthInstanceID: core.StringPtr("testString"),
			}

			response, err := watsonxDataService.DeletePrestoEngineCatalogs(deletePrestoEngineCatalogsOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(204))
		})
	})

	Describe(`DeleteSparkEngine - Delete spark engine`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`DeleteSparkEngine(deleteSparkEngineOptions *DeleteSparkEngineOptions)`, func() {
			deleteSparkEngineOptions := &watsonxdatav2.DeleteSparkEngineOptions{
				EngineID: core.StringPtr("testString"),
				AuthInstanceID: core.StringPtr("testString"),
			}

			response, err := watsonxDataService.DeleteSparkEngine(deleteSparkEngineOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(204))
		})
	})

	Describe(`DeleteSparkEngineApplications - Stop Spark Applications`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`DeleteSparkEngineApplications(deleteSparkEngineApplicationsOptions *DeleteSparkEngineApplicationsOptions)`, func() {
			deleteSparkEngineApplicationsOptions := &watsonxdatav2.DeleteSparkEngineApplicationsOptions{
				EngineID: core.StringPtr("testString"),
				ApplicationID: core.StringPtr("testString"),
				AuthInstanceID: core.StringPtr("testString"),
				State: []string{"testString"},
			}

			response, err := watsonxDataService.DeleteSparkEngineApplications(deleteSparkEngineApplicationsOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(204))
		})
	})

	Describe(`DeleteSparkEngineCatalogs - Disassociate catalogs from a spark engine`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`DeleteSparkEngineCatalogs(deleteSparkEngineCatalogsOptions *DeleteSparkEngineCatalogsOptions)`, func() {
			deleteSparkEngineCatalogsOptions := &watsonxdatav2.DeleteSparkEngineCatalogsOptions{
				EngineID: core.StringPtr("testString"),
				CatalogNames: core.StringPtr("testString"),
				AuthInstanceID: core.StringPtr("testString"),
			}

			response, err := watsonxDataService.DeleteSparkEngineCatalogs(deleteSparkEngineCatalogsOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(204))
		})
	})

	Describe(`DeleteSparkEngineHistoryServer - Stop spark history server`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`DeleteSparkEngineHistoryServer(deleteSparkEngineHistoryServerOptions *DeleteSparkEngineHistoryServerOptions)`, func() {
			deleteSparkEngineHistoryServerOptions := &watsonxdatav2.DeleteSparkEngineHistoryServerOptions{
				EngineID: core.StringPtr("testString"),
				AuthInstanceID: core.StringPtr("testString"),
			}

			response, err := watsonxDataService.DeleteSparkEngineHistoryServer(deleteSparkEngineHistoryServerOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(204))
		})
	})

	Describe(`DeleteSchema - Delete schema`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`DeleteSchema(deleteSchemaOptions *DeleteSchemaOptions)`, func() {
			deleteSchemaOptions := &watsonxdatav2.DeleteSchemaOptions{
				EngineID: core.StringPtr("testString"),
				CatalogID: core.StringPtr("testString"),
				SchemaID: core.StringPtr("testString"),
				AuthInstanceID: core.StringPtr("testString"),
			}

			response, err := watsonxDataService.DeleteSchema(deleteSchemaOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(204))
		})
	})

	Describe(`DeleteTable - Delete table`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`DeleteTable(deleteTableOptions *DeleteTableOptions)`, func() {
			deleteTableOptions := &watsonxdatav2.DeleteTableOptions{
				CatalogID: core.StringPtr("testString"),
				SchemaID: core.StringPtr("testString"),
				TableID: core.StringPtr("testString"),
				EngineID: core.StringPtr("testString"),
				AuthInstanceID: core.StringPtr("testString"),
			}

			response, err := watsonxDataService.DeleteTable(deleteTableOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(204))
		})
	})

	Describe(`DeleteColumn - Delete column`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`DeleteColumn(deleteColumnOptions *DeleteColumnOptions)`, func() {
			deleteColumnOptions := &watsonxdatav2.DeleteColumnOptions{
				EngineID: core.StringPtr("testString"),
				CatalogID: core.StringPtr("testString"),
				SchemaID: core.StringPtr("testString"),
				TableID: core.StringPtr("testString"),
				ColumnID: core.StringPtr("testString"),
				AuthInstanceID: core.StringPtr("testString"),
			}

			response, err := watsonxDataService.DeleteColumn(deleteColumnOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(204))
		})
	})

	Describe(`DeleteMilvusService - Delete milvus service`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`DeleteMilvusService(deleteMilvusServiceOptions *DeleteMilvusServiceOptions)`, func() {
			deleteMilvusServiceOptions := &watsonxdatav2.DeleteMilvusServiceOptions{
				ServiceID: core.StringPtr("testString"),
				AuthInstanceID: core.StringPtr("testString"),
			}

			response, err := watsonxDataService.DeleteMilvusService(deleteMilvusServiceOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(204))
		})
	})

	Describe(`DeleteIngestionJobs - Delete an ingestion job`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`DeleteIngestionJobs(deleteIngestionJobsOptions *DeleteIngestionJobsOptions)`, func() {
			deleteIngestionJobsOptions := &watsonxdatav2.DeleteIngestionJobsOptions{
				JobID: core.StringPtr("testString"),
				AuthInstanceID: core.StringPtr("testString"),
			}

			response, err := watsonxDataService.DeleteIngestionJobs(deleteIngestionJobsOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(204))
		})
	})
})

//
// Utility functions are declared in the unit test file
//
