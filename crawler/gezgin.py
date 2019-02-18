import os, time, json
from selenium import webdriver

PROJECT_ROOT = os.path.abspath(os.path.dirname(__file__))
DRIVER_BIN = os.path.join(PROJECT_ROOT, "driver/chromedriver")

base_url = 'https://www.avmgezgini.com/avmler/'

driver = webdriver.Chrome(executable_path = DRIVER_BIN)
driver2 = webdriver.Chrome(executable_path = DRIVER_BIN)
driver.get(base_url)
driver2.get(base_url)

boxes = driver.find_element_by_class_name('avm_liste').find_elements_by_class_name('avm_kutu')
avmLines = []
for box in boxes:

    name = box.find_element_by_tag_name('p').text
    score = box.find_element_by_class_name('avm_puan').find_element_by_tag_name('b').text
    img_src = box.find_element_by_tag_name('img').get_attribute('src')
    detail_page = box.find_element_by_tag_name('a').get_attribute('href')

    print('Crawling ===> ' + name)

    data = {}
    data['ImageLink'] = str(img_src)
    data['DetailPage'] = str(detail_page)
    data['AvmName'] = str(name)
    data['AvmScore'] = str(score)

    # print("Image link  : " + str(img_src))
    # print("Detail page : " + str(detail_page))
    # print("Avm name    : " + str(name))
    # print("Avm score   : " + str(score))
    # print()

    # Avm Detail Page
    driver2.get(detail_page)
    magazalar_link = driver2.find_element_by_xpath('//*[@id="EsitSolAlan"]/div[3]/a[1]').get_attribute('href')
    cafe_rest_link = driver2.find_element_by_xpath('//*[@id="EsitSolAlan"]/div[3]/a[2]').get_attribute('href')
    
    data['Shops'] = {}

    # Avm List of Shops
    driver2.get(magazalar_link)
    dukkans = driver2.find_elements_by_xpath('//*[@id="EsitSolAlan"]/div[5]/table/tbody/tr')
    shopId = 0
    for dukkan in dukkans:
        logo = dukkan.find_element_by_tag_name('img').get_attribute('src')
        magaza = dukkan.find_element_by_xpath('td[2]').text
        kat = dukkan.find_element_by_xpath('td[3]').text
        telefon = dukkan.find_element_by_xpath('td[4]').text

        data['Shops'][shopId] = {}
        data['Shops'][shopId]['Logo'] = str(logo)
        data['Shops'][shopId]['Magaza'] = str(magaza)
        data['Shops'][shopId]['Kat'] = str(kat)
        data['Shops'][shopId]['Telefon'] = str(telefon)
        shopId = shopId + 1

    driver2.back()

    data['Cafes'] = {}
    # Avm List of Cafes
    driver2.get(cafe_rest_link)
    dukkans = driver2.find_elements_by_xpath('//*[@id="EsitSolAlan"]/div[5]/table/tbody/tr')
    cafeId = 0
    for dukkan in dukkans:
        logo = dukkan.find_element_by_tag_name('img').get_attribute('src')
        magaza = dukkan.find_element_by_xpath('td[2]').text
        kat = dukkan.find_element_by_xpath('td[3]').text
        telefon = dukkan.find_element_by_xpath('td[4]').text

        data['Cafes'][cafeId] = {}
        data['Cafes'][cafeId]['Logo'] = str(logo)
        data['Cafes'][cafeId]['Magaza'] = str(magaza)
        data['Cafes'][cafeId]['Kat'] = str(kat)
        data['Cafes'][cafeId]['Telefon'] = str(telefon)
        cafeId = cafeId + 1
    avmLines.append(data)

with open('data.txt', 'a') as outfile:
    for hostDict in avmLines:
        json.dump(hostDict, outfile)
        outfile.write('\n')

time.sleep(5)
driver.close()