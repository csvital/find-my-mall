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
   

    print('Started crawling ' + name)

    data = {}
    data['ImageLink'] = str(img_src)
    data['DetailPage'] = str(detail_page)
    data['Name'] = str(name)
    data['Score'] = str(score)

    # Avm Detail Page
    driver2.get(detail_page)
    magazalar_link = driver2.find_element_by_xpath('//*[@id="EsitSolAlan"]/div[3]/a[1]').get_attribute('href')
    cafe_rest_link = driver2.find_element_by_xpath('//*[@id="EsitSolAlan"]/div[3]/a[2]').get_attribute('href')
    address = driver2.find_element_by_class_name('avmd_adres').text

    words = str(address).split()
    district = ' '.join(words[:-1])
    city = words[-1]

    data['District'] = str(district)
    data['City'] = str(city)
    
    data['Shops'] = []

    print('Crawling shops of ', name)
    # Avm List of Shops
    driver2.get(magazalar_link)
    dukkans = driver2.find_elements_by_xpath('//*[@id="EsitSolAlan"]/div[5]/table/tbody/tr')
    for dukkan in dukkans:
        logo = dukkan.find_element_by_tag_name('img').get_attribute('src')
        magaza = dukkan.find_element_by_xpath('td[2]').text
        kat = dukkan.find_element_by_xpath('td[3]').text
        telefon = dukkan.find_element_by_xpath('td[4]').text
        data['Shops'].append({'Logo': str(logo), 'Magaza': str(magaza), 'Kat': str(kat), 'Telefon': str(telefon)})

    driver2.back()

    print('Crawling cafes of ', name)
    data['Cafes'] = []
    # Avm List of Cafes
    driver2.get(cafe_rest_link)
    dukkans = driver2.find_elements_by_xpath('//*[@id="EsitSolAlan"]/div[5]/table/tbody/tr')
    for dukkan in dukkans:
        logo = dukkan.find_element_by_tag_name('img').get_attribute('src')
        magaza = dukkan.find_element_by_xpath('td[2]').text
        kat = dukkan.find_element_by_xpath('td[3]').text
        telefon = dukkan.find_element_by_xpath('td[4]').text
        data['Cafes'].append({'Logo': str(logo), 'Magaza': str(magaza), 'Kat': str(kat), 'Telefon': str(telefon)})

    avmLines.append(data)

filename = '../data/data.txt'
print('Started to writing into ', filename)
with open(filename, 'a') as outfile:
    for hostDict in avmLines:
        json.dump(hostDict, outfile)
        outfile.write('\n')
print('Finished writing')

time.sleep(5)
driver.close()