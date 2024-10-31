async function makeSoapRequestPUT(event) {
    event.preventDefault(); // Prevent default form submission

    const formData = new FormData(event.target); // Get form data
    const profile = {
        Address: {
            AddressId: formData.get('addressId'), // Ensure you include this in your form
            City: formData.get('city'),
            Country: formData.get('country'),
            PostalCode: formData.get('postalCode'),
            State: formData.get('state'),
            Street: formData.get('street'),
        },
        Birthday: formData.get('birthday'),
        Gender: formData.get('gender'),
        Nationality: formData.get('nationality'),
        PhoneNumber: formData.get('phoneNumber'),
        SocialSecurityNumber: formData.get('socialSecurityNumber'),
    };

    const soapEnvelope = `
        <soap:Envelope xmlns:soap="http://schemas.xmlsoap.org/soap/envelope/" xmlns:tns="http://tempuri.org/">
            <soap:Body>
                <tns:UpdateProfile>
                    <tns:profile>
                        <tns:Address>
                            <tns:AddressId>${profile.Address.AddressId}</tns:AddressId>
                            <tns:City>${profile.Address.City}</tns:City>
                            <tns:Country>${profile.Address.Country}</tns:Country>
                            <tns:PostalCode>${profile.Address.PostalCode}</tns:PostalCode>
                            <tns:State>${profile.Address.State}</tns:State>
                            <tns:Street>${profile.Address.Street}</tns:Street>
                        </tns:Address>
                        <tns:Birthday>${profile.Birthday}</tns:Birthday>
                        <tns:Gender>${profile.Gender}</tns:Gender>
                        <tns:Nationality>${profile.Nationality}</tns:Nationality>
                        <tns:PhoneNumber>${profile.PhoneNumber}</tns:PhoneNumber>
                        <tns:SocialSecurityNumber>${profile.SocialSecurityNumber}</tns:SocialSecurityNumber>
                    </tns:profile>
                </tns:UpdateProfile>
            </soap:Body>
        </soap:Envelope>
    `;

    try {
        const response = await fetch('http://localhost:8002/ProfileService.asmx', {
            method: 'POST',
            headers: {
                'Content-Type': 'text/xml; charset=utf-8',
                'SOAPAction': 'http://tempuri.org/IProfileService/UpdateProfile'
            },
            body: soapEnvelope
        });

        const textResponse = await response.text();
        console.log(textResponse);
        
        // Optionally, handle the response here
        // e.g., parse the response and give feedback to the user

    } catch (error) {
        console.error('Error making SOAP request:', error);
        // Optionally, display an error message to the user
    }
}
