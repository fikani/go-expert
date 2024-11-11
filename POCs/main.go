package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

// Configuration for the database connection.
// Update these variables with your actual database credentials.
const (
	dbUser     = ""
	dbPassword = ""
	dbHost     = "master-aurora-cluster.cluster-cl2shdkfnpqv.us-east-1.rds.amazonaws.com"
	dbPort     = "3306"
	dbName     = "soci"
)

// Query to calculate the average response time for a given account_id.
// The notif_type list is hardcoded as per your requirements.
const avgResponseTimeQuery = `
SELECT
    AVG(
        IF(
            resolved_by != 0,
            TIMESTAMPDIFF(SECOND, timestamp, resolved_at),
            NULL
        )
    ) AS avg_response_time_seconds
FROM
    tasks
WHERE
    timestamp BETWEEN '2024-06-01' AND '2024-08-31'
    AND notif_type IN (
        'NOTIF_TYPE_FB_PM',
        'NOTIF_TYPE_FB_COMMENT',
        'NOTIF_TYPE_FB_MSG',
        'NOTIF_TYPE_FB_MENTION_POST',
        'NOTIF_TYPE_FB_MENTION_COMMENT',
        'NOTIF_TYPE_TW_MENTION',
        'NOTIF_TYPE_TW_REPLY',
        'NOTIF_TYPE_TW_QUOTE',
        'NOTIF_TYPE_TW_PM',
        'NOTIF_TYPE_GP_COMMENT',
        'NOTIF_TYPE_LI_COMMENT',
        'NOTIF_TYPE_IG_COMMENT',
        'NOTIF_TYPE_IG_BUSINESS_COMMENT',
        'NOTIF_TYPE_IG_BUSINESS_MENTION_MEDIA',
        'NOTIF_TYPE_IG_BUSINESS_MENTION_COMMENT',
        'NOTIF_TYPE_IG_BUSINESS_MEDIA_TAG',
        'NOTIF_TYPE_IG_BUSINESS_DIRECT_MESSAGE',
        'NOTIF_TYPE_GMB_QUESTION',
        'NOTIF_TYPE_GMB_ANSWER',
        'NOTIF_TYPE_GMB_GOOGLE_BUSINESS_MESSAGES',
        'NOTIF_TYPE_SMS_MESSAGE',
        'NOTIF_TYPE_CHAT_WIDGET'
    )
    AND account_id = ?
;`

func main() {
	// List of account_ids as provided.
	accountIDs := []int{
		731, 758, 777, 870, 1046, 1071, 1075, 1085, 1102, 1113, 1123, 1134, 1139, 1142, 1150, 1153, 1160, 1164, 1175, 1190, 1192, 1211, 1232, 1233, 1237, 1266, 1294, 1312, 1331, 1332, 1352, 1393, 1399, 1401, 1416, 1420, 1421, 1424, 1427, 1446, 1449, 1481, 1482, 1490, 1496, 1498, 1527, 1535, 1536, 1545, 1567, 1569, 1576, 1582, 1583, 1584, 1596, 1601, 1615, 1621, 1632, 1634, 1656, 1673, 1678, 3000, 3003, 3004, 3007, 3015, 3029, 3045, 3051, 3053, 3056, 3075, 3082, 3089, 3094, 3097, 3101, 3108, 3117, 3133, 3136, 3139, 3140, 3141, 3147, 3150, 3153, 3154, 3157, 3158, 3159, 3160, 3170, 3174, 3180, 3189, 3193, 3196, 3197, 3201, 3202, 3211, 3219, 3236, 3237, 3247, 3248, 3254, 3262, 3263, 3270, 3286, 3288, 3295, 3296, 3300, 3302, 3303, 3332, 3351, 3363, 3364, 3367, 3368, 3370, 3373, 3374, 3375, 3380, 3384, 3385, 3403, 3404, 3405, 3453, 3462, 3469, 3476, 3478, 3482, 3485, 3486, 3491, 3495, 3507, 3508, 3516, 3527, 3528, 3537, 3546, 3548, 3554, 3557, 3563, 3578, 3589, 3603, 3604, 3612, 3619, 3622, 3643, 3649, 3651, 3656, 3658, 3659, 3660, 3661, 3662, 3664, 3666, 3670, 3675, 3682, 3686, 3687, 3690, 3691, 3702, 3706, 3711, 3728, 3736, 3744, 3745, 3747, 3753, 3768, 3780, 3812, 3817, 3826, 3830, 3831, 3834, 3839, 3849, 3852, 3867, 3882, 3883, 3888, 3901, 3903, 3936, 3937, 3948, 3951, 3958, 3967, 3991, 3994, 3996, 4005, 4009, 4016, 4018, 4031, 4033, 4041, 4045, 4046, 4061, 4068, 4084, 4091, 4094, 4098, 4117, 4121, 4124, 4183, 4195, 4215, 4219, 4220, 4235, 4239, 4240, 4241, 4260, 4261, 4270, 4277, 4283, 4304, 4321, 4326, 4334, 4347, 4350, 4351, 4356, 4360, 4363, 4366, 4369, 4370, 4376, 4378, 4451, 4459, 4480, 4501, 4515, 4652, 4664, 4676, 4681, 4685, 4691, 4692, 4714, 4724, 4742, 4757, 4758, 4774, 4779, 4785, 4825, 4867, 5152, 5216, 5382, 5414, 5415, 5450, 5458, 5481, 5502, 5529, 5532, 5563, 5564, 5605, 5617, 5635, 5639, 5647, 5654, 5655, 5690, 5719, 5743, 5756, 5757, 5760, 5761, 5779, 5781, 5792, 5823, 5861, 5870, 5905, 5908, 5910, 5922, 5991, 5994, 6009, 6013, 6037, 6056, 6058, 6061, 6066, 6069, 6089, 6091, 6106, 6118, 6120, 6121, 6122, 6123, 6124, 6125, 6126, 6127, 6128, 6129, 6130, 6131, 6132, 6133, 6134, 6135, 6136, 6137, 6138, 6139, 6140, 6141, 6142, 6144, 6145, 6146, 6147, 6148, 6149, 6150, 6155, 6156, 6157, 6158, 6159, 6160, 6161, 6162, 6163, 6164, 6165, 6166, 6167, 6168, 6169, 6171, 6172, 6174, 6175, 6176, 6177, 6178, 6179, 6180, 6181, 6182, 6183, 6184, 6185, 6186, 6187, 6188, 6189, 6190, 6191, 6192, 6193, 6194, 6195, 6196, 6197, 6198, 6200, 6201, 6202, 6203, 6204, 6205, 6206, 6207, 6208, 6209, 6210, 6211, 6212, 6213, 6214, 6215, 6216, 6217, 6218, 6219, 6220, 6221, 6222, 6223, 6224, 6225, 6226, 6227, 6228, 6229, 6232, 6233, 6234, 6237, 6238, 6239, 6240, 6241, 6242, 6243, 6244, 6245, 6246, 6247, 6248, 6249, 6250, 6257, 6267, 6286, 6305, 6314, 6333, 6338, 6361, 6376, 6379, 6387, 6423, 6473, 6474, 6479, 6482, 6483, 6488, 6489, 6492, 6537, 6580, 6581, 6584, 6585, 6610, 6630, 6635, 6644, 6645, 6648, 6669, 6688, 6697, 6700, 6713, 6719, 6728, 6748, 6788, 6812, 6822, 6843, 6844, 6862, 6865, 6916, 6920, 6921, 6924, 6933, 6947, 6951, 6966, 6967, 6975, 6995, 7015, 7035, 7053, 7070, 7084, 7096, 7105, 7115, 7116, 7121, 7122, 7132, 7133, 7136, 7180, 7181, 7182, 7194, 7196, 7214, 7215, 7235, 7266, 7267, 7317, 7318,
		7358,
	}

	// Database connection string.
	// Format: username:password@tcp(host:port)/dbname
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=true", dbUser, dbPassword, dbHost, dbPort, dbName)

	// Open the database connection.
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		log.Fatalf("Error opening database: %v", err)
	}
	defer db.Close()

	// Verify the connection to the database.
	if err := db.Ping(); err != nil {
		log.Fatalf("Error connecting to database: %v", err)
	}
	fmt.Println("Successfully connected to the database.")

	// Prepare the statement for executing the average response time query.
	stmt, err := db.Prepare(avgResponseTimeQuery)
	if err != nil {
		log.Fatalf("Error preparing statement: %v", err)
	}
	defer stmt.Close()

	// Map to store the average response time per account_id.
	accountAverages := make(map[int]float64)

	// Variables to compute the general average.
	var totalResponseTime float64
	var totalResolvedTasks int

	// Iterate through each account_id sequentially.
	for _, accountID := range accountIDs {
		var avg sql.NullFloat64

		// Execute the query with the current account_id.
		err := stmt.QueryRow(accountID).Scan(&avg)
		if err != nil {
			log.Printf("Error querying account_id %d: %v", accountID, err)
			continue
		}

		// If the average is valid, store it and update the total response time and task count.
		if avg.Valid {
			accountAverages[accountID] = avg.Float64
			totalResponseTime += avg.Float64
			totalResolvedTasks++
			fmt.Printf("Account ID %d: Average Response Time = %.2f seconds\n", accountID, avg.Float64)
		} else {
			fmt.Printf("Account ID %d: No resolved tasks found.\n", accountID)
		}
	}

	// Check if there are any resolved tasks.
	if totalResolvedTasks == 0 {
		log.Println("No resolved tasks found for the specified accounts and date range.")
		return
	}

	// Compute the general average response time.
	generalAvgResponseTime := totalResponseTime / float64(totalResolvedTasks)

	// Display the general average.
	fmt.Printf("\nGeneral Average Response Time Across All Accounts: %.2f seconds\n", generalAvgResponseTime)
}
