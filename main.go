package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// 備忘錄架構
type Todolist struct {
	ID   int
	ToDo string
	Done bool
}

// 使用者架構
type User struct {
	Name     string
	Todolist []Todolist
	NextID   int
}

// 編號控制架構
type MenuItem struct {
	Number int    // 編號
	Name   string // 選項名稱
	Action string // 關聯操作
}

func main() {
	var userList []User
	reader := bufio.NewReader(os.Stdin) //用於讀取標準輸入的bufio.Reader物件。
	innerMenu := true

MainLoop: //整體循環
	for innerMenu {
		fmt.Print("User name: ")
		username, _ := reader.ReadString('\n')
		username = strings.TrimSpace(username)

		// 檢查是否已存在該使用者
		var currentUser *User
		for i := range userList {
			if userList[i].Name == username {
				currentUser = &userList[i]
				break
			}
		}

		// 如果使用者不存在，則創建一個新的使用者
		if currentUser == nil {
			user := User{Name: username}
			userList = append(userList, user)
			currentUser = &userList[len(userList)-1]
		}

	InnerMenu: //主菜單內部循環
		for {
			// 定義菜單選項
			menuItems := []MenuItem{
				{Number: 1, Name: "新增備忘錄明細(add)", Action: "add"},
				{Number: 2, Name: "檢視備忘錄清單(list)", Action: "list"},
				{Number: 3, Name: "查看使用者列表(list users)", Action: "list users"},
				{Number: 4, Name: "退出備忘錄(exit)", Action: "exit"},
			}

			// 打印菜單選項
			fmt.Println("選擇操作 : ")
			for _, item := range menuItems {
				fmt.Printf("[%d]%s ", item.Number, item.Name)
			}
			fmt.Print("\n輸入選擇: ")

			// 讀取用戶輸入的選擇
			choiceInput, _ := reader.ReadString('\n')
			choiceInput = strings.TrimSpace(choiceInput)

			// 將用戶輸入的選擇轉為整數
			choiceNum, err := strconv.Atoi(choiceInput)
			if err != nil {
				fmt.Println("無效的選擇. 請輸入有效的選擇編號.")
				continue
			}

			// 根據用戶的選擇執行相應的操作
			var selectedAction string
			for _, item := range menuItems {
				if choiceNum == item.Number {
					selectedAction = item.Action
					break
				}
			}

			switch selectedAction {
			case "add": //進入 add 子菜單

				fmt.Print("Enter todo and done (e.g., Buy groceries true/false): ")
				input, _ := reader.ReadString('\n')
				input = strings.TrimSpace(input)
				var done bool

				parts := strings.Fields(input)
				todo := strings.Join(parts[:len(parts)-1], " ") //讓使用者輸入內容更為彈性

				if len(parts) < 2 {
					fmt.Println("Invalid input. Please enter at least todo and done values.")
					continue
				}

				if strings.ToLower(parts[len(parts)-1]) == "true" {
					done = true
				} else if strings.ToLower(parts[len(parts)-1]) == "false" {
					done = false
				} else {
					fmt.Println("Invalid input for 'done'. Please enter 'true' or 'false'.")
					continue
				}

				newItem := Todolist{
					ID:   currentUser.NextID,
					ToDo: todo,
					Done: done,
				}

				currentUser.Todolist = append(currentUser.Todolist, newItem)
				currentUser.NextID++
				fmt.Println("已新增至備忘錄.")

				// 等待用户按 Enter 键
				fmt.Print("按 Enter 鍵繼續...")
				reader.ReadString('\n')

			// 进入 list 子菜单
			case "list":

				if len(currentUser.Todolist) == 0 {
					fmt.Println("備忘錄清單為空，請添加.")

					// 等待用戶按 Enter 鍵
					fmt.Print("按 Enter 鍵繼續...")
					reader.ReadString('\n')
					continue //如果沒有備忘錄項目，直接返回主菜單
				}

				//建立內部循環
				for {
					fmt.Println("備忘錄清單 : ")
					for i, todo := range currentUser.Todolist {
						fmt.Printf("[%d] num:%d , todo:%s, done:%v\n", i+1, todo.ID, todo.ToDo, todo.Done)
					}

					// 顯示備忘錄項目操作選項
					fmt.Println("-------------------------")
					fmt.Println("請選擇備忘錄項目操作:")
					fmt.Printf("[a] 編輯備忘錄項目(edit), [b] 刪除備忘錄項目(delete), [c] 返回主選單 : ")
					editOrDeleteInput, _ := reader.ReadString('\n')
					editOrDeleteInput = strings.TrimSpace(editOrDeleteInput)

					switch editOrDeleteInput {
					case "a":
						// 要求用戶輸入要編輯項目編號
						fmt.Print("輸入要編輯的備忘錄項目的編號 (或 [c] 返回上一頁) : ")
						editInput, _ := reader.ReadString('\n')
						editInput = strings.TrimSpace(editInput)

						if editInput != "c" {
							// 將用戶的輸入轉為整數
							editNum, err := strconv.Atoi(editInput)
							if err != nil {
								fmt.Println("無效的輸入，請輸入有效的編號或 [c] 返回上一頁.")
								fmt.Print("按 Enter 鍵繼續...")
								reader.ReadString('\n')
								continue
							}
							if editNum > 0 && editNum <= len(currentUser.Todolist) {
								// 编辑選定的備忘錄
								indexToEdit := editNum - 1
								fmt.Printf("目前的項目: %s (完成狀態: %v)\n", currentUser.Todolist[indexToEdit].ToDo, currentUser.Todolist[indexToEdit].Done)

								// 詢問用戶是否要編輯狀態
								fmt.Println("要編輯什麼? (1) 備忘錄項目, (2) 完成度狀態 : ")
								editOptionInput, _ := reader.ReadString('\n')
								editOptionInput = strings.TrimSpace(editOptionInput)

								switch editOptionInput {
								case "1":
									// 詢問用戶要設置的新項目
									fmt.Print("輸入新的備忘錄項目 : ")
									newTodo, _ := reader.ReadString('\n')
									newTodo = strings.TrimSpace(newTodo)
									currentUser.Todolist[indexToEdit].ToDo = newTodo
									fmt.Println("備忘錄項目已成功編輯.")
								case "2":
									// 詢問用戶要設置的新狀態
									fmt.Print("輸入新的狀態 (true/false): ")
									newStatusInput, _ := reader.ReadString('\n')
									newStatusInput = strings.TrimSpace(newStatusInput)

									// 將用戶輸入轉為布林值
									newStatus, err := strconv.ParseBool(newStatusInput)
									if err != nil {
										fmt.Println("無效的狀態，請輸入 'true' 或 'false'.")
										fmt.Print("按 Enter 鍵繼續...")
										reader.ReadString('\n')
										continue
									}

									// 更新代辦事項狀況
									currentUser.Todolist[indexToEdit].Done = newStatus
									fmt.Println("備忘錄項目的狀態已成功編輯.")
								default:
									fmt.Println("無效的選擇，請輸入 '1' 或 '2'.")
									fmt.Print("按 Enter 鍵繼續...")
									reader.ReadString('\n')
								}
								fmt.Print("按 Enter 鍵繼續...")
								reader.ReadString('\n')
							} else {
								fmt.Println("無效的編號，請輸入有效的編號.")
								fmt.Print("按 Enter 鍵繼續...")
								reader.ReadString('\n')
							}
						} else {
							// 返回主菜单
							break
						}

					case "b":
						// 要求用户输入要删除的项目编号
						fmt.Print("輸入要刪除的備忘錄項目的編號或 (或 [c] 返回上一頁) : ")
						deleteInput, _ := reader.ReadString('\n')
						deleteInput = strings.TrimSpace(deleteInput)

						if deleteInput != "c" {
							// 将用户的输入转换为整数
							deleteNum, err := strconv.Atoi(deleteInput)
							if err != nil {
								fmt.Println("無效的輸入，請輸入有效的編號.")
								continue
							} else if deleteNum > 0 && deleteNum <= len(currentUser.Todolist) {
								// 删除选定的备忘录项目
								indexToDelete := deleteNum - 1
								currentUser.Todolist = append(currentUser.Todolist[:indexToDelete], currentUser.Todolist[indexToDelete+1:]...)
								fmt.Println("備忘錄項目已成功刪除.")
								fmt.Print("按 Enter 鍵繼續...")
								reader.ReadString('\n')
							} else {
								fmt.Println("無效的編號，請輸入有效的編號.")
								fmt.Print("按 Enter 鍵繼續...")
								reader.ReadString('\n')
							}
						} else {
							// 返回主菜单
							break
						}
					case "c":
						// 返回主菜单
						continue InnerMenu
					}
				}

			case "list users":
				//使用布林值來判斷是否回到innermenu 還是繼續內部循環
				backToInnerMenu := false
				for !backToInnerMenu {
					fmt.Println("使用者清單 : ")
					for i, user := range userList {
						fmt.Printf("[%d] num:%d , user:%s\n", i+1, user.NextID, user.Name)
					}

					// 顯示分隔线
					fmt.Println("-------------------------")
					// 顯示当前用户信息
					fmt.Printf("當前使用者: %s\n", currentUser.Name)

					fmt.Print("輸入 [a] 切換使用者(change),[b] 刪除使用者(delete),[c] 創建新使用者(new),[d] 返回備忘錄(back) : ")
					changeUserInput, _ := reader.ReadString('\n')
					changeUserInput = strings.TrimSpace(changeUserInput)

					switch changeUserInput {
					case "a":
						if len(userList) == 1 {
							fmt.Println("未找到其他使用者，請 [c] 建立新使用者.")
							fmt.Print("按 Enter 鍵繼續...")
							reader.ReadString('\n')
							continue
						}

						// 要求用户输入要切换的用户編號
						fmt.Print("輸入要切換的使用者的編號 (或 [d] 返回上一頁) : ")
						changeUserNumStr, _ := reader.ReadString('\n')
						changeUserNumStr = strings.TrimSpace(changeUserNumStr)

						if changeUserNumStr != "d" {
							// 將用户的输入轉換為整数
							changeUserNum, err := strconv.Atoi(changeUserNumStr)
							if err != nil {
								fmt.Println("無效的編號，請輸入有效的編號.")
								fmt.Print("按 Enter 鍵繼續...")
								reader.ReadString('\n')
								continue
							}

							if changeUserNum > 0 && changeUserNum <= len(userList) {
								indexToChange := changeUserNum - 1
								currentUser = &userList[indexToChange]
								fmt.Printf("已切換到使用者 %s.\n", currentUser.Name)
							} else {
								fmt.Println("無效的編號，請輸入有效的編號.")
							}
						} else {
							// 用户取消切换，返回到主菜单
							fmt.Print("按 Enter 鍵繼續...")
							reader.ReadString('\n')
							backToInnerMenu = true
						}

					case "b":
						if len(userList) == 1 {
							fmt.Println("未找到其他使用者，請 [c] 建立新使用者.")
							fmt.Print("按 Enter 鍵繼續...")
							reader.ReadString('\n')
							continue
						}

						// 要求用户输入要删除的用户编號
						fmt.Print("輸入要刪除的使用者的編號 (或 [d] 返回上一頁) : ")
						deleteUserNumStr, _ := reader.ReadString('\n')
						deleteUserNumStr = strings.TrimSpace(deleteUserNumStr)

						if deleteUserNumStr != "d" {
							// 將用戶的輸入轉為整數
							deleteUserNum, err := strconv.Atoi(deleteUserNumStr)
							if err != nil {
								fmt.Println("無效的編號，請輸入有效的編號.")
								fmt.Print("按 Enter 鍵繼續...")
								reader.ReadString('\n')
								continue
							}

							if deleteUserNum > 0 && deleteUserNum <= len(userList) {
								indexToDelete := deleteUserNum - 1
								fmt.Printf("已刪除使用者 %s.\n", userList[indexToDelete].Name)
								userList = append(userList[:indexToDelete], userList[indexToDelete+1:]...)
								//退出當前循環，返回到主菜單
								backToInnerMenu = true
							} else {
								fmt.Println("無效的編號，請輸入有效的編號.")
							}
						} else {
							// 用户取消删除，返回到主菜单
							backToInnerMenu = true
						}

					case "c":
						fmt.Print("輸入新使用者名稱 : ")
						newUserName, _ := reader.ReadString('\n')
						newUserName = strings.TrimSpace(newUserName)

						var newUser *User
						for i := range userList {
							if userList[i].Name == newUserName {
								newUser = &userList[i]
								fmt.Print("按 Enter 鍵繼續...")
								reader.ReadString('\n')
								break
							}
						}

						if newUser != nil {
							fmt.Println("使用者已存在.")
						} else {
							user := User{Name: newUserName}
							userList = append(userList, user)
							newUser = &userList[len(userList)-1]
							fmt.Printf("已創建新使用者 %s.\n", newUser.Name)

						}

					case "d":
						backToInnerMenu = true
					}
				}

			case "exit":
				fmt.Println("拜拜!")
				innerMenu = false
				break MainLoop // 退出主循環

			default:
				fmt.Println("無效指令. 請輸入有效的選擇編號.")
			}
		}
	}
}
